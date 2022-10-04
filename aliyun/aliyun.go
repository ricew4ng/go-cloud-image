package aliyun

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

/*
	aliyun openapi client
*/

type IAliyunClient interface {
	// 获取所有region列表
	DescribeRegions() ([]Region, error)
	// 获取region对应的所有image
	DescribeImages(regionID string) (images []Image, err error)
	// 共享镜像
	ShareImage(regionID, imageID, accountID string) error
	// 复制镜像
	CopyImage(imageID, srcRegion, dstRegion string) error
}

func NewAliyunClient(ak, sk string) (IAliyunClient, error) {
	credential := credentials.NewAccessKeyCredential(ak, sk)
	cli, err := ecs.NewClientWithOptions(defaultRegionID,
		sdk.NewConfig(), credential)
	if err != nil {
		return nil, err
	}
	return &aliyunClient{
		//credential: credential,
		aliyunCli: cli,
	}, nil
}

type aliyunClient struct {
	//credential *credentials.AccessKeyCredential
	aliyunCli *ecs.Client
}

const (
	defaultRegionID = "cn-hangzhou"
)

func (c *aliyunClient) ShareImage(regionID, imageID, accountID string) error {
	request := ecs.CreateModifyImageSharePermissionRequest()
	request.Scheme = "https"
	request.ImageId = imageID
	request.RegionId = regionID
	request.AddAccount = &[]string{accountID}
	response, err := c.aliyunCli.ModifyImageSharePermission(request)
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return errors.New(response.String())
	}
	return nil
}

func (c *aliyunClient) CopyImage(srcRegionID, srcImageID, dstRegionID string) error {
	request := ecs.CreateCopyImageRequest()
	request.Scheme = "https"
	request.ImageId = srcImageID
	request.RegionId = srcRegionID
	request.DestinationRegionId = dstRegionID
	response, err := c.aliyunCli.CopyImage(request)
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return errors.New(response.String())
	}
	return nil
}

func (c *aliyunClient) DescribeRegions() (regions []Region, err error) {
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := c.aliyunCli.DescribeRegions(request)
	if err != nil {
		return regions, err
	}
	if !response.IsSuccess() {
		return regions, errors.New(response.String())
	}
	for _, r := range response.Regions.Region {
		regions = append(regions, Region{
			Label:    r.LocalName,
			EndPoint: r.RegionEndpoint,
			RegionID: r.RegionId,
		})
	}
	return regions, nil
}

func (c *aliyunClient) DescribeImages(regionID string) (images []Image, err error) {
	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.RegionId = regionID
	request.IsPublic = requests.NewBoolean(false)
	request.ImageOwnerAlias = "self"
	request.PageSize = requests.NewInteger(100)
	response, err := c.aliyunCli.DescribeImages(request)
	if err != nil {
		return images, err
	}
	if !response.IsSuccess() {
		return images, errors.New(response.String())
	}
	for _, img := range response.Images.Image {
		images = append(images, Image{
			Label:   img.ImageName,
			ImageID: img.ImageId,
			OS:      img.OSName,
			Status:  img.Status,
		})
	}
	return images, nil
}
