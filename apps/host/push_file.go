package host

import (
	"context"
	"fmt"
	"github.com/pkg/sftp"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// 打开本地文件
func openLocalFile(path string) (*FileInfo, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, err
	}

	fileObj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// 文件信息，内容，权限、大小
	f := &FileInfo{
		Content: fileObj,
		Model:   stat.Mode(),
		Size:    stat.Size(),
	}

	return f, nil
}

// 打开远端文件
func (c *ClientConfig) openRemoteFile(path string, mode os.FileMode) (*sftp.File, error) {
	// 1 判断文件是否存在,如果不存在则先创建文件存放目录
	_, err := c.SFTPClient.Stat(path)
	if os.IsNotExist(err) {
		_, file := filepath.Split(path)
		err := c.SFTPClient.MkdirAll(file)
		if err != nil {
			return nil, err
		}
	}

	// 2 创建文件
	fileObj, err := c.SFTPClient.Create(path)
	if err != nil {
		return nil, err
	}

	// 3 设置远程文件的权限
	if err := fileObj.Chmod(mode); err != nil {
		return nil, err
	}

	return fileObj, nil
}

// DistributeFile 文件下发
func DistributeFile(config *ClientConfig, src *FileInfo, dst string, wg *sync.WaitGroup) {
	defer wg.Done()

	dstFile, err := config.openRemoteFile(dst, src.Model)
	if err != nil {
		return
	}

	_, err = io.Copy(dstFile, src.Content)
	if err != nil {
		return
	}

}

func InitSFTP(hostList []*ClientConfig, src, dst string) {
	var wg sync.WaitGroup

	// 获取本地文件的信息（文件内容、大小、权限）
	srcFileInfo, err := openLocalFile(src)
	if err != nil {
		return
	}

	// 循环主机列表进行文件下发
	for _, v := range hostList {
		err := v.CreateClientForSecretKey(context.Background())
		if err != nil {
			fmt.Printf("host %s: %s\n", v.Host, err)
			continue
		}

		sftpClient, err := sftp.NewClient(v.Client)
		if err != nil {

			continue
		}
		v.SFTPClient = sftpClient

		wg.Add(1)

		go DistributeFile(v, srcFileInfo, dst, &wg)
	}

	// 等待所有goroutine完成
	wg.Wait()
}