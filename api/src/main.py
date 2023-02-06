from http import HTTPStatus

import aioredis
from elasticsearch import AsyncElasticsearch
from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from fastapi.responses import ORJSONResponse
from starlette.responses import PlainTextResponse

from api.v1 import films, genres, persons
from core.config import settings
from infrastructure import elastic, redis

app = FastAPI(
    title=settings.project_name,
    docs_url='/api/openapi',
    openapi_url='/api/openapi.json',
    default_response_class=ORJSONResponse,
)


@app.on_event('startup')
async def startup():
    redis.redis = await aioredis.create_redis_pool(
        (settings.storage_host, settings.storage_port),
        minsize=settings.redis_pool_min_size,
        maxsize=settings.redis_pool_max_size,
    )
    elastic.es = AsyncElasticsearch(hosts=['{host}:{port}'.format(
        host=settings.es_host,
        port=settings.es_port,
    )])


@app.on_event('shutdown')
async def shutdown():
    await redis.redis.close()
    await elastic.es.close()


@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request, exc):
    return PlainTextResponse(str(exc), status_code=HTTPStatus.BAD_REQUEST)


app.include_router(films.router, prefix='/api/v1/films', tags=['films'])
app.include_router(genres.router, prefix='/api/v1/genres', tags=['genres'])
app.include_router(persons.router, prefix='/api/v1/persons', tags=['persons'])
