package blog

import (
	"strconv"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase"
)

type BlogUcStruct struct {
	repo repo.BlogRepoInterface
	log  *log.LogCustom
}

func NewBlogUc(repo repo.BlogRepoInterface, log *log.LogCustom) usecase.BlogUcInterface {
	return &BlogUcStruct{
		repo: repo,
		log:  log,
	}
}

func (b BlogUcStruct) AddBlog(v models.Blog) (models.Blog, error) {
	return b.repo.AddBlog(v)
}

func (b BlogUcStruct) GetAll() ([]models.Blog, error) {
	return b.repo.GetAll()
}

func (b BlogUcStruct) GetById(id int) (models.Blog, error) {
	return b.repo.GetById(id)
}

func (b BlogUcStruct) UpdateData(id int, v models.Blog) (models.Blog, error) {
	return b.repo.UpdateData(id, v)
}

func (b BlogUcStruct) DeleteData(id []string) error {
	for _, s := range id {
		idRes, _ := strconv.Atoi(s)
		_, err := b.repo.GetById(idRes)
		if err != nil {
			b.log.Error(err, "usecase error when get data by id", "", nil, idRes, nil)
			return err
		}
	}

	return b.repo.DeleteData(id)
}
