package utils

import (
	"bytes"
	"context"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"net/http"
)

// ReadRemoteExcelFile Get the file according to the CDN link
func ReadRemoteExcelFile(fileUrl string) (file *excelize.File, err error) {
	resp, err := http.Get(fileUrl)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return
	}
	if resp == nil || resp.Body == nil {
		return
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	file, err = excelize.OpenReader(bytes.NewReader(buf))
	if err != nil {
		return
	}
	return
}

// GetExcelRows Get each line of data in the file according to the CDN link
func GetExcelRows(ctx context.Context, url string) ([][]string, error) {
	fileHandler, err := ReadRemoteExcelFile(url)
	if err != nil {
		return nil, err
	}
	sheetIndex := fileHandler.GetActiveSheetIndex()
	if sheetIndex == 0 {
		return nil, errors.New("excel格式不正确")
	}
	rows := fileHandler.GetRows(fileHandler.GetSheetName(sheetIndex))
	return rows, nil
}
