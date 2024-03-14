package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hellpgf/internal/consts"
	"hellpgf/internal/dao"
	"hellpgf/internal/model"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// Upload 定义图片上传位置
//校验：上传位置是否正确，安全性校验，一分钟之内只能上传十次
//定义年月日 Ymd
//入库
//返回数据

func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	//定义图片上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败上传路径不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	//安全性校验:每个人一分钟内只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	//避免在代码中写死常量
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("传入频繁，一分钟内只能上传十次")
	}
	//定义年月日
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	//gfile.Join用反斜线拼接
	if err != nil {
		return nil, err
	}
	data := entity.FileInfo{Name: fileName,
		Src: gfile.Join(uploadPath, dateDirName, fileName),
		Url: "/upload/" + dateDirName + "/" + fileName, //和上面gfile.Join效果一样
	}
	//入库   OmitEmpty()过滤字段
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
