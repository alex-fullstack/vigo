from http import HTTPStatus

from fastapi import APIRouter, Depends, HTTPException


router = APIRouter()


@router.get('/', response_model=list[FileShort])
async def files(
    file_service: FileService = Depends(get_file_service),
) -> list[FileShort]:
    """
    Список загруженных файлов.

    Parameters:
        genre_service: сервис загруженных файлов

    Returns:
        list[FileShort]: страница загруженыых файлов с заданными параметрами
    """
    return await file_service.get_list()


@router.get('/{file_id}', response_model=File)
async def file_details(
    file_id: str,
    file_service: FileService = Depends(get_file_service),
) -> File:
    """
    Детальная страница файла.

    Parameters:
        file_service: сервис загруженных файлов
        file_id: идентификатор записи

    Raises:
        HTTPException: исключение со статусом 404

    Returns:
        File: страница с данными загруженного файла
    """
    file = await file_service.get_by_id(file_id=file_id)
    if not file:
        raise HTTPException(
            status_code=HTTPStatus.NOT_FOUND,
            detail='file not found',
        )

    return file
