package render

import (
	"bbs-go/model"

	"bbs-go/util/simple"
	"bbs-go/util/simple/json"

	"github.com/sirupsen/logrus"
)

func buildImageList(imageListStr string) (imageList []model.ImageInfo) {
	if simple.IsNotBlank(imageListStr) {
		var images []model.ImageDTO
		if err := json.Parse(imageListStr, &images); err == nil {
			if len(images) > 0 {
				for _, image := range images {
					imageList = append(imageList, model.ImageInfo{
						Url:     HandleOssImageStyleDetail(image.Url),
						Preview: HandleOssImageStylePreview(image.Url),
					})
				}
			}
		} else {
			logrus.Error(err)
		}
	}
	return
}
