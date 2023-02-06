from typing import Type

from pydantic import BaseModel

from abstract_repository import AbstractEntityRepository


class BaseEntityService:
    def __init__(
        self,
        repository: AbstractEntityRepository,
    ):
        self.repository = repository

    async def get(self, *, entity_id: str, scheme: Type[BaseModel]) -> BaseModel | None:
        entity = await self.repository.get(entity_id=entity_id)
        if not entity:
            return None
        response = scheme(**entity)
        return response
