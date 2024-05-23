package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/hopehook/golang-db/mysql"
	"kpi/internal/dictionary"
	"kpi/internal/dto"
	"kpi/internal/entity"
)

type DbPool struct {
	db *mysql.SQLConnPool
}

var DB *DbPool

// GetDb создает и возвращает пул подключений к БД
func GetDb() *DbPool {
	if DB == nil {
		d := dictionary.Env
		DB = &DbPool{
			db: mysql.InitMySQLPool(d.DbHost, d.DbName, d.DbUser, d.DbPassword, "utf8", 3, 1),
		}

	}
	return DB
}

// saveFact сохраняет fact в бд
func (d *DbPool) saveFact(f *entity.Fact) (id int64, err error) {
	query, args, err := sq.
		Insert("fact").
		Columns("period_start", "period_end", "period_key", "indicator_to_mo_id", "indicator_to_mo_fact_id", "value", "fact_time", "is_plan", "auth_user_id", "comment").
		Values(f.PeriodStart, f.PeriodEnd, f.PeriodKey, f.IndicatorToMoId, f.IndicatorToMoFactId, f.Value, f.FactTime, f.IsPlan, f.AuthUserId, f.Comment).
		ToSql()
	if err != nil {
		return 0, err
	}

	return d.db.Insert(query, args...)
}

// GetFacts получает список facts
func (d *DbPool) GetFacts(factsReq *dto.GetFactsRequest) ([]map[string]interface{}, error) {
	query, args, err := sq.
		Select("id", "period_start", "period_end", "period_key", "indicator_to_mo_id", "indicator_to_mo_fact_id", "value", "fact_time", "is_plan", "auth_user_id", "comment").
		From("fact").
		Where(sq.And{sq.Eq{"period_start": factsReq.PeriodStart}, sq.Eq{"period_end": factsReq.PeriodEnd}, sq.Eq{"period_key": factsReq.PeriodKey}, sq.Eq{"indicator_to_mo_id": factsReq.IndicatorToMoId}}).
		ToSql()
	if err != nil {
		return nil, err
	}
	return d.db.Query(query, args...)
}
