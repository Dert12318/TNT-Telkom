package expedition_schedule_uc

import (
	"CartApp/config/env"
	"CartApp/config/log"
	"CartApp/models"
	"CartApp/models/contract"
	"CartApp/repo"
	"CartApp/usecase"
	"CartApp/utils"
	"errors"
	"strconv"
	"strings"
	"time"
)

type EsUcStruct struct {
	repo repo.ExpeditionRepoInterface
	log  *log.LogCustom
}

func NewEsUc(repo repo.ExpeditionRepoInterface, log *log.LogCustom) usecase.ExpeditionUcInterface {
	return &EsUcStruct{
		repo: repo,
		log:  log,
	}
}

func (e EsUcStruct) AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	var tempRoute string
	arr := []string{v.RouteFrom, v.RouteDestination}
	for _, val := range arr {
		atoi, err := strconv.Atoi(val)
		if err != nil {
			return models.ExpeditionSchedule{}, errors.New(contract.ErrBadRequest)
		}
		res, err := e.repo.GetKotaById(atoi)
		if err != nil {
			return models.ExpeditionSchedule{}, err
		}
		tempRoute += res.Nama + "-"
	}

	routeArr := strings.Split(tempRoute, "-")
	v.Route = routeArr[0] + "-" + routeArr[1]

	if !utils.CheckDate(v.Eta, v.Etd, v.Closing) {
		return models.ExpeditionSchedule{}, errors.New(contract.ErrBadRequest)
	}

	layoutISO := "2006-01-02"
	closingTemp, err := time.Parse(layoutISO, v.Closing)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}
	etdTemp, err := time.Parse(layoutISO, v.Etd)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}
	etaTemp, err := time.Parse(layoutISO, v.Eta)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	if closingTemp.Unix() > etdTemp.Unix() || etdTemp.Unix() > etaTemp.Unix() || closingTemp.Unix() > etaTemp.Unix() {
		return models.ExpeditionSchedule{}, errors.New(contract.ErrDateSchedule)
	}

	es, err := e.repo.AddEs(v)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	return es, nil
}

func (e EsUcStruct) GetById(id int) (models.ExpeditionSchedule, error) {
	es, err := e.repo.GetById(id)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	return es, nil
}

func (e EsUcStruct) GetAll(limit, offset, uri string) (models.ResponseDataPagination, error) {

	limitTemp, _ := strconv.Atoi(limit)
	offsetTemp, _ := strconv.Atoi(offset)

	if limitTemp == 0 {
		limitTemp = 1
	}

	if offsetTemp <= 0 {
		offsetTemp = 1
	}

	all, err := e.repo.GetAll(limitTemp, offsetTemp)
	if err != nil {
		return models.ResponseDataPagination{}, err
	}

	for i, _ := range all.Data {
		all.Data[i].RouteFrom = ""
		all.Data[i].RouteDestination = ""

		if len(all.Data[i].Eta) != 10 || len(all.Data[i].Etd) != 10 || len(all.Data[i].Closing) != 10 {
			return models.ResponseDataPagination{}, errors.New(contract.ErrBadRequest)
		}
		monthEta := utils.ConvertMonth(all.Data[i].Eta)
		etaTemp := strings.Split(all.Data[i].Eta, "-")
		str := etaTemp[2] + "-" + monthEta + "-" + etaTemp[0]
		etaTempArr := strings.Split(str, "-")
		eta := strings.Join(etaTempArr, " ")
		all.Data[i].Eta = eta

		monthEtd := utils.ConvertMonth(all.Data[i].Etd)
		etdTemp := strings.Split(all.Data[i].Etd, "-")
		strEtd := etdTemp[2] + "-" + monthEtd + "-" + etdTemp[0]
		etdTempArr := strings.Split(strEtd, "-")
		etd := strings.Join(etdTempArr, " ")
		all.Data[i].Etd = etd

		monthClosing := utils.ConvertMonth(all.Data[i].Closing)
		closingTemp := strings.Split(all.Data[i].Closing, "-")
		strClosing := closingTemp[2] + "-" + monthClosing + "-" + closingTemp[0]
		cosingTempArr := strings.Split(strClosing, "-")
		closing := strings.Join(cosingTempArr, " ")
		all.Data[i].Closing = closing
		//etdTemp := strings.Split(all[i].Etd, "-")
		//closingTemp := strings.Split(all[i].Closing, "-")
	}

	lengthData := all.TotalData
	eachData := all.NumberEnd

	if lengthData == 0 {
		lengthData = limitTemp
	}

	if limitTemp > lengthData {
		limitTemp = lengthData
	}

	tempTotalPage := lengthData / limitTemp

	if lengthData%limitTemp != 0 {
		tempTotalPage = tempTotalPage + 1
	}

	all.TotalData = lengthData
	all.DataPerPage = limitTemp
	all.TotalPage = tempTotalPage

	if eachData < limitTemp {
		all.NumberEnd = (limitTemp * (offsetTemp - 1)) + eachData
	} else {
		all.NumberEnd = limitTemp * offsetTemp
	}

	if offsetTemp == 1 {
		all.NumberCurrent = 1
	} else if all.TotalPage == offsetTemp && eachData == 1 {
		all.NumberCurrent = all.NumberEnd
	} else if all.TotalPage == offsetTemp && eachData != 1 {
		all.NumberCurrent = (all.NumberEnd - eachData) + 1
	} else {
		all.NumberCurrent = (all.NumberEnd - limitTemp) + 1
	}

	all.Page = offsetTemp

	for i2, _ := range all.Data {
		all.Data[i2].No = i2 + all.NumberCurrent
	}

	tempUrl := strings.Split(uri, "&page=")
	atoi, _ := strconv.Atoi(tempUrl[1])
	tempPageNext := atoi + 1
	tempPagePrev := atoi - 1
	if tempPagePrev < 1 {
		tempPagePrev = 1
	}
	nextPage := strconv.Itoa(tempPageNext)
	prevPage := strconv.Itoa(tempPagePrev)

	if all.Page <= 1 {
		if all.DataPerPage == all.TotalData {
			all.PrevUrlPage = ""
			all.NextUrlPage = ""
		} else {
			all.PrevUrlPage = ""
			all.NextUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + nextPage
		}
	} else if all.Page == all.TotalPage {
		all.PrevUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + prevPage
		all.NextUrlPage = ""
	} else {
		all.PrevUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + prevPage
		all.NextUrlPage = env.Config.Protocol + "://" + tempUrl[0] + "&page=" + nextPage
	}

	return all, nil
}

func (e EsUcStruct) Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (e EsUcStruct) DeleteData(id []string) error {
	for _, s := range id {
		idRes, _ := strconv.Atoi(s)
		_, err := e.repo.GetById(idRes)
		if err != nil {
			e.log.Error(err, "usecase error when get data by id", "", nil, idRes, nil)
			return err
		}
	}

	err := e.repo.DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}

func (e EsUcStruct) GetByRoute(route string) ([]models.ExpeditionSchedule, error) {
	routeTemp := strings.ToUpper(route)
	routeArr := strings.Split(routeTemp, "-")
	es, err := e.repo.GetByRoute(routeArr[0], routeArr[1])
	if err != nil {
		return []models.ExpeditionSchedule{}, err
	}

	for i, _ := range es {
		es[i].RouteFrom = ""
		es[i].RouteDestination = ""

		monthEta := utils.ConvertMonth(es[i].Eta)
		etaTemp := strings.Split(es[i].Eta, "-")
		str := etaTemp[2] + "-" + monthEta + "-" + etaTemp[0]
		etaTempArr := strings.Split(str, "-")
		eta := strings.Join(etaTempArr, " ")
		es[i].Eta = eta

		monthEtd := utils.ConvertMonth(es[i].Etd)
		etdTemp := strings.Split(es[i].Etd, "-")
		strEtd := etdTemp[2] + "-" + monthEtd + "-" + etdTemp[0]
		etdTempArr := strings.Split(strEtd, "-")
		etd := strings.Join(etdTempArr, " ")
		es[i].Etd = etd

		monthClosing := utils.ConvertMonth(es[i].Closing)
		closingTemp := strings.Split(es[i].Closing, "-")
		strClosing := closingTemp[2] + "-" + monthClosing + "-" + closingTemp[0]
		cosingTempArr := strings.Split(strClosing, "-")
		closing := strings.Join(cosingTempArr, " ")
		es[i].Closing = closing
	}

	return es, nil
}
