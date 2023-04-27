package repository

import (
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/entity"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *entity.Course) error
	Update(id int, updatedCourse *entity.Course) error
	Delete(course *entity.Course) error
	FindById(id int) (*entity.Course, error)
	FindBySlug(slug string) (*entity.Course, error)
}

type courseRepositoryImpl struct {
	db *gorm.DB
}

type CourseRConfig struct {
	DB *gorm.DB
}

func NewCourseRepository(c *CourseRConfig) CourseRepository {
	return &courseRepositoryImpl{
		db: c.DB,
	}
}

func(r *courseRepositoryImpl) Create(course *entity.Course) error {
	err := r.db.Create(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *courseRepositoryImpl) Update(id int, updatedCourse *entity.Course) error {
	var course entity.Course
	r.db.First(&course, id)
	err := r.db.Model(&course).Updates(updatedCourse).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *courseRepositoryImpl) Delete(course *entity.Course) error {
	err := r.db.Delete(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *courseRepositoryImpl) FindById(id int) (course *entity.Course, err error) {
	err = r.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}

func(r *courseRepositoryImpl) FindBySlug(slug string) (course *entity.Course, err error) {
	err = r.db.Where("slug = ?", slug).First(&course).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}