package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type UserCourseRepository interface {
	Insert(*entity.UserCourse) error
	Update(*entity.UserCourse) error
	Delete(Id int) error
	FindAll(userId int) ([]*entity.UserCourse, error)
	Count(courseId int) (int64, error)
	Find(userId int, courseId int) (*entity.UserCourse, error)
}

type usercourseRepositoryImpl struct {
	db *gorm.DB
}

type UserCourseRConfig struct {
	DB *gorm.DB
}

func NewUserCourseRepository(c *UserCourseRConfig) UserCourseRepository {
	return &usercourseRepositoryImpl{
		db: c.DB,
	}
}

func (r *usercourseRepositoryImpl) Insert(userCourse *entity.UserCourse) error {
	err := r.db.Create(&userCourse).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *usercourseRepositoryImpl) Update(updatedUserCourse *entity.UserCourse) error {
	var userCourse entity.UserCourse
	r.db.Where("course_id = ?", updatedUserCourse.CourseId).Where("user_id", updatedUserCourse.UserId).First(&userCourse)
	err := r.db.Model(&userCourse).Updates(updatedUserCourse).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *usercourseRepositoryImpl) Delete(Id int) error {
	err := r.db.Delete(&entity.UserCourse{}, Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *usercourseRepositoryImpl) FindAll(userId int) (usercourses []*entity.UserCourse, err error) {
	err = r.db.Preload("Course").Where("user_id = ?", userId).Find(&usercourses).Error
	if err != nil {
		return nil, err
	}
	return usercourses, nil
}

func (r *usercourseRepositoryImpl) Count(CourseId int) (total int64, err error) {
	err = r.db.Model(&entity.UserCourse{}).Where("course_id = ?", CourseId).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *usercourseRepositoryImpl) Find(userId int, courseId int) (usercourse *entity.UserCourse, err error) {
	r.db.Unscoped().Where("course_id = ?", courseId).Where("user_id = ?", userId).First(&usercourse)
	if err != nil {
		return nil, err
	}
	return usercourse, nil
}
