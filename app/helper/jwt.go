package helper

//
//type (
//	Area struct {
//		ID        int64   `gorm:"column:id;primaryKey;"`
//		AreaValue float64 `gorm:"column:area_value"`
//		AreaType  string  `gorm:"column:type"`
//	}
//)
//
//func (ar Area) save(value float64, areaType string) error {
//	ar.AreaValue = value
//	ar.AreaType = areaType
//	if err := _r.DB.create(&ar).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (_r *AreaRepository) InsertArea(param1 float64, param2 float64, typeArea string) (err error) {
//	ar := Area{}
//	switch typeArea {
//	case "persegi panjang":
//		if err := ar.save((param1 * param2), typeArea); err != nil {
//			return err
//		}
//	case "persegi":
//		if err := ar.save((param1 * param2), typeArea); err != nil {
//			return err
//		}
//	case "segitiga":
//		area := 0.5 * (param1 * param2)
//		if err := ar.save(area, typeArea); err != nil {
//			return err
//		}
//	}
//	return nil
//}

//
//type

//
//func (a *Area) setArea(value int64,areatype string)  {
//	a.AreaValue = value
//	a.AreaType = areatype
//}
//func (Area) save()  {
//	err = _r.DB.create(&Area{}).Error
//	if err != nil {
//		return err
//	}
//}
//type AreaRepository interface {
//	Area() float64
//}
//
//type persegi struct {
//	//Area float64
//}
//
//func (p persegi) Area() float64 {
//	return 2
//}
//
//type persegi struct {
//	//Area float64
//}
//
//func (p persegi) Area() float64 {
//	return 2
//}
//
////func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area)
////(err error) {
////
////}
