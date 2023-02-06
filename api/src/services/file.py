from functools import lru_cache

from fastapi import Depends
from services.base import BaseEntityService


class FileService(BaseEntityService):
    async def get_list(self) -> list:
        return await self.repository.get_list(page_size=1000, page_number=1)

    async def get_by_id(self, *, file_id: str) -> File | None:
        return await self.get(entity_id=file_id, scheme=File)


@lru_cache()
def get_file_service(
    storage: Any = Depends(get_redis),
    task_repo: elasticsearch.AsyncElasticsearch = Depends(get_elastic),
    processing_repo: Any = Depends(),
) -> FileService:
    return FileService()
