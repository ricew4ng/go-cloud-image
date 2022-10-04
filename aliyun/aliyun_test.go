package aliyun

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewAliyunClient(t *testing.T) {
	Convey("Test Aliyun Client", t, func() {
		ak := `<your-ak>`
		sk := `<your-sk>`
		cli, err := NewAliyunClient(ak, sk)
		So(err, ShouldBeNil)
		Convey("Test Get Regions", func() {
			regions, err := cli.DescribeRegions()
			So(err, ShouldBeNil)
			for _, r := range regions {
				t.Log(r.Label, r.EndPoint, r.RegionID)
			}
		})
		Convey("Test Get Images", func() {
			regionID := "cn-hangzhou"
			imgs, err := cli.DescribeImages(regionID)
			So(err, ShouldBeNil)
			for _, i := range imgs {
				t.Log(i.Label, i.ImageID, i.OS, i.Status)
			}
		})
		Convey("Test Share Image", func() {
			regionID := `cn-hangzhou`
			imageID := ``
			accountID := ``
			err := cli.ShareImage(regionID, imageID, accountID)
			So(err, ShouldBeNil)
		})
		Convey("Test Copy Image", func() {
			srcImageID := ``
			srcRegionID := `cn-hangzhou`
			dstRegionID := `cn-qingdao`
			err := cli.CopyImage(srcRegionID, srcImageID, dstRegionID)
			So(err, ShouldBeNil)
		})
	})
}
