package common

type Pagination struct {
	Page     int // Trang hiện tại (bắt đầu từ 1)
	PageSize int // Số lượng bản ghi mỗi trang
}

func (p Pagination) Offset() int {
	if p.Page < 1 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

func (p Pagination) Limit() int {
	if p.PageSize <= 0 {
		return 10 // default
	}
	return p.PageSize
}
