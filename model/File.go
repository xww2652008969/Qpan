package model

type File struct {
	ID          uint `gorm:"primaryKey"`
	Useruuid    string
	name        string
	Md5         string
	Fileaddress string
}

//func (file File) Getfile(m map[string]interface{}) (File, bool) {
//	var cfile File
//	service.Condb.Where(m).First(&cfile)
//	if service.Condb.Error == nil {
//		return cfile, true
//	} else {
//		return cfile, false
//	}
//}
//
//// 修改user表 如果成功则返回ture和影响行数,否则返回false，和影响行数
//func (file File) Updatefale(m map[string]interface{}, o map[string]string) (bool, int64) {
//	ref := service.Condb.Model(&file).Where(m).Update(o["key"], o["value"])
//	fmt.Println(ref.RowsAffected)
//	if service.Condb.Error == nil {
//		return true, ref.RowsAffected
//	} else {
//		return false, ref.RowsAffected
//	}
//}
