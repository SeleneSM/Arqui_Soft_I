package services

import (
	Clients "Arqui_Soft_I/clients/archivos"
	Dto "Arqui_Soft_I/dto"
	Model "Arqui_Soft_I/model"
	e "Arqui_Soft_I/utils"
)

func CreateFile(fileDto Dto.ArchivoDto, fileContent []byte) (Model.Archivo, e.ApiError) {
	file := Model.Archivo{
		Name:     fileDto.Name,
		Curso_Id: fileDto.Curso_id,
	}

	filePath, err := Clients.SaveFile(fileDto.Name, fileContent)
	if err != nil {
		return file, err
	}

	file.Path = filePath
	return Clients.CreateFile(file)
}
