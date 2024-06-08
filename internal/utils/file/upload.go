package appfile

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/httprespond"
)

func Init() {
	_ = os.MkdirAll(getUploadTempPath(), 0600)
}

func getUploadTempPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "files/temp")
}

func GetFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", getUploadTempPath(), fileName)
}

func RemoveFile(fileName string) {
	_ = os.Remove(GetFilePath(fileName))
}

func UploadSingle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = c.Get("ctx").(*appcontext.AppContext)

		file, ok := c.Get("file").(*multipart.FileHeader)
		if !ok {
			ctx.Logger().ErrorText("file not found")
			return httprespond.R400(c, apperrors.Common.InvalidFile, nil)
		}

		src, err := file.Open()
		if err != nil {
			ctx.Logger().Error("failed to open file", err, appcontext.Fields{"file": file.Filename})
			return httprespond.R400(c, apperrors.Common.InvalidFile, nil)
		}
		defer func() { _ = src.Close() }()

		dst, err := os.Create(path.Join(getUploadTempPath(), file.Filename))
		if err != nil {
			ctx.Logger().Error("failed to create file to upload temporary directory", err, appcontext.Fields{"file": file.Filename})
			return httprespond.R400(c, apperrors.Common.InvalidFile, nil)
		}
		defer func() { _ = dst.Close() }()

		if _, err = io.Copy(dst, src); err != nil {
			ctx.Logger().Error("failed to copy file to upload temporary directory", err, appcontext.Fields{"file": file.Filename})
			return httprespond.R400(c, apperrors.Common.InvalidFile, nil)
		}

		return next(c)
	}
}
