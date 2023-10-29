/*
 * @Date: 2023-10-25 05:59:02
 * @LastEditors: hxlh
 * @LastEditTime: 2023-10-28 12:33:03
 * @FilePath: /1024/server/src/storage/video_dao_impl.go
 */
package storage

import (
	"database/sql"
	"dev1024/src/entities"

	"github.com/pkg/errors"
)

type VideoDaoImpl struct {
	baseDB *BaseDB
}

func NewVideoDaoImpl(baseDB *BaseDB) *VideoDaoImpl {
	return &VideoDaoImpl{
		baseDB: baseDB,
	}
}

func (t *VideoDaoImpl) GetNextNByVid(vid int64, n int) ([]*entities.VideoInfo, error) {
	tx, err := t.baseDB.DB.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ans, err := t.getNextNByVid(tx, vid, n)
	return ans, commitOrRollback(err, tx)
}

func (t *VideoDaoImpl) getNextNByVid(tx *sql.Tx, vid int64, n int) ([]*entities.VideoInfo, error) {
	stmt, err := tx.Prepare("SELECT * FROM video1024.video_info WHERE vid > ? LIMIT 0,?")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(vid, n)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	
	ans := make([]*entities.VideoInfo, 0)
	for rows.Next() {
		videoInfo := &entities.VideoInfo{}
		rows.Scan(&videoInfo.Vid, &videoInfo.UpLoader, &videoInfo.CDN, &videoInfo.Subtitled, &videoInfo.Likes, &videoInfo.Tags)
		ans = append(ans, videoInfo)
	}
	return ans, nil
}

func (t *VideoDaoImpl) Save(videoInfo *entities.VideoInfo) error {
	tx, err := t.baseDB.DB.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	err = t.save(tx, videoInfo)
	return commitOrRollback(err, tx)
}

func (t*VideoDaoImpl) save(tx *sql.Tx, videoInfo *entities.VideoInfo) error {
	stmt, err := tx.Prepare("INSERT INTO video1024.video_info(uploader,cdn,subtitled,likes,tags) VALUES(?,?,?,?,?)")
	if err != nil {
		return errors.WithStack(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(videoInfo.UpLoader, videoInfo.CDN, videoInfo.Subtitled, videoInfo.Likes, videoInfo.Tags)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

var _ VideoDao = (*VideoDaoImpl)(nil)

