package data

import "go-web-example/models"

var courses = []*models.Course{
	{
		ID:            "1",
		Title:         "Course 1",
		Author:        "Author 1",
		PublishedDate: "2021-01-01",
	},
}

func ListCourses() []*models.Course {
	return courses
}

func GetCourse(ID string) *models.Course {
	for _, course := range courses {
		if course.ID == ID {
			return course
		}
	}
	return nil
}

func CreateCourse(course models.Course) {
	courses = append(courses, &course)
}

func UpdateCourse(ID string, course models.Course) *models.Course {
	for i, c := range courses {
		if c.ID == ID {
			courses[i] = &course
			return c
		}
	}
	return nil
}

func DeleteCourse(ID string) *models.Course {
	for i, c := range courses {
		if c.ID == ID {
			courses = append(courses[:i], courses[i+1:]...)
			return &models.Course{}
		}
	}
	return nil
}
