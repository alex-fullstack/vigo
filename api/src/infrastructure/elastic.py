from typing import Optional

import elasticsearch
from elasticsearch import AsyncElasticsearch
from fastapi.exceptions import RequestValidationError
from pydantic.error_wrappers import ErrorWrapper

from core.config import QUERY_LITERAL, HITS_LITERAL
from services.abstract_repository import AbstractEntityRepository

es: Optional[AsyncElasticsearch] = None


class ElasticSearchMixin:
    def __init__(self, *, db: AsyncElasticsearch):
        self.db = db

    async def _get(
        self,
        *,
        entity_id: str,
        index: str,
    ) -> dict | None:
        query = {'term': {'id': {'value': entity_id}}}
        entities = await self._search(query=query, index=index, offset=0, size=1)
        return entities[0] if entities else None

    async def _search(
        self,
        *,
        query: dict = None,
        sort: str = '',
        offset: int,
        size: int,
        index: str,
    ) -> list[dict]:
        sort_direction = 'desc' if sort.startswith('-') else 'asc'
        request_params = {
            'size': size,
            'sort': [{sort.removeprefix('-'): sort_direction}] if sort else [],
            'from': offset,
        }
        if query:
            request_params[QUERY_LITERAL] = query
        try:
            search = await self.db.search(request_params, index)
            return [
                hit['_source'] for hit in search[HITS_LITERAL][HITS_LITERAL]
            ]
        except elasticsearch.exceptions.RequestError as error:
            raise RequestValidationError(
                errors=[ErrorWrapper(error, error.error)],
            )


class ElasticSearchGenreRepository(AbstractEntityRepository, ElasticSearchMixin):
    async def get(self, *, entity_id) -> dict | None:
        return await self._get(entity_id=entity_id, index='genres')

    async def get_list(
        self,
        *,
        page_size: int,
        page_number: int,
        sort: str = '',
        filtration: dict = None,
    ) -> list[dict]:
        return await self._search(
            index='genres',
            offset=page_size * (page_number - 1),
            size=page_size,
        )


class ElasticSearchPersonRepository(AbstractEntityRepository, ElasticSearchMixin):
    async def get(self, *, entity_id) -> dict | None:
        return await self._get(entity_id=entity_id, index='persons')

    async def get_list(
            self,
            *,
            page_size: int,
            page_number: int,
            query: dict = None,
            sort: str = '',
    ) -> list[dict]:
        return await self._search(
            query=query,
            sort=sort,
            index='persons',
            offset=page_size * (page_number - 1),
            size=page_size,
        )


class ElasticSearchFilmRepository(AbstractEntityRepository, ElasticSearchMixin):
    async def get(self, *, entity_id) -> dict | None:
        return await self._get(entity_id=entity_id, index='movies')

    async def get_list(
        self,
        *,
        page_size: int,
        page_number: int,
        query: dict = None,
        sort: str = '',
    ) -> list[dict]:
        return await self._search(
            query=query,
            sort=sort,
            index='movies',
            offset=page_size * (page_number - 1),
            size=page_size,
        )


# Функция понадобится при внедрении зависимостей
async def get_elastic() -> AsyncElasticsearch:
    return es
