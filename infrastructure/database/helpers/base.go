package dbHelpers

import (
	"database/sql"
	"fmt"
	"time"

	sharedDTO "git.techpartners.asia/gateway-services/payment-service/internal/shared/dto"
	utilsPkg "git.techpartners.asia/gateway-services/payment-service/pkg/utils"
	"gorm.io/gorm"
)

type DBHelper struct {
	*gorm.DB
}

func NewOrm(db *gorm.DB) *DBHelper {
	return &DBHelper{DB: db}
}

func (d *DBHelper) Paginate(input *sharedDTO.SharedPaginationRequestDTO) *DBHelper {
	d.DB = d.DB.Scopes(func(d *gorm.DB) *gorm.DB {
		if input == nil {
			return d
		}

		if input.Limit == 0 {
			input.Limit = 20
		}
		return d.Offset((input.Page) * input.Limit).Limit(input.Limit)
	})

	return d
}

func (d *DBHelper) Entity(entity interface{}) *DBHelper {
	d.DB = d.DB.Model(&entity)
	return d
}

// join
func (d *DBHelper) Join(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("INNER JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) LeftJoin(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("LEFT JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) RightJoin(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("RIGHT JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) Search(fields []string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	queryString := ""
	for index, field := range fields {
		if index == 0 {
			queryString = fmt.Sprintf("LOWER(%s) like LOWER(@value)", field)
			continue
		}
		fieldQuery := fmt.Sprintf("LOWER(%s) like LOWER(@value)", field)
		queryString = fmt.Sprintf("%s or %s", queryString, fieldQuery)
	}
	d.DB = d.DB.Where(queryString, sql.Named("value", "%"+*value+"%"))
	return d
}

func (d *DBHelper) In(field string, value interface{}) *DBHelper {
	if value == nil || utilsPkg.IsNil(value) {
		return d
	}
	queryString := fmt.Sprintf("%s in (?)", field)
	d.DB = d.DB.Where(queryString, value)
	return d
}

func (d *DBHelper) Total() (int64, error) {
	var total int64
	if err := d.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (d *DBHelper) Equal(field string, value interface{}) *DBHelper {
	if utilsPkg.IsNil(value) {
		return d
	}

	d.DB = d.DB.Where(fmt.Sprintf("%s = ?", field), value)
	return d
}

func (d *DBHelper) NotEqual(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s != ?", field), *value)
	return d
}

func (d *DBHelper) NotIn(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s not in (?)", field), *value)
	return d
}

func (d *DBHelper) Bool(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s = ?", field), *value == "1")
	return d
}

func (d *DBHelper) Like(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s like ?", field), "%"+*value+"%")
	return d
}

func (d *DBHelper) GreaterDate(field string, value *time.Time) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s >= ?", field), value)
	return d
}

func (d *DBHelper) LessDate(field string, value *time.Time) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s <= ?", field), value)
	return d
}

func (d *DBHelper) GreaterYear(field string, value *time.Time) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("extract(year from %s) >= ?", field), value.Year())
	return d
}

func (d *DBHelper) LessYear(field string, value *time.Time) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("extract(year from %s) <= ?", field), value.Year())
	return d
}

func (d *DBHelper) BetweenDates(field string, minDate *time.Time, maxDate *time.Time) *DBHelper {
	if minDate == nil || maxDate == nil {
		return d
	}

	d.DB = d.DB.Where(fmt.Sprintf("date(%s) between date(?) and date(?)", field), minDate, maxDate)
	return d
}

func (d *DBHelper) BetweenDateTimes(field string, times []*string) *DBHelper {
	if len(times) < 2 || len(times) > 2 {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s between ? and ?", field), times[0], times[1])
	return d
}
