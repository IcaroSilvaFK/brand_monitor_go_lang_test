package configs

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

type CloudinaryConfig struct {
	Cld *cloudinary.Cloudinary
}

func NewCloudinaryConfig() (*CloudinaryConfig, error) {

	cld, err := cloudinary.NewFromURL("cloudinary://511675566119594:Hfmf5siUSIUYnFflLPK8E079xt0@dlf01tbs6")

	if err != nil {
		return nil, err
	}

	return &CloudinaryConfig{
		Cld: cld,
	}, nil
}
