from logging import config as logging_config

from pydantic import BaseSettings

from core.logger import LOGGING

QUERY_LITERAL = 'query'
HITS_LITERAL = 'hits'


class Settings(BaseSettings):
    project_name: str
    storage_host: str
    storage_port: str
    redis_pool_min_size: int
    redis_pool_max_size: int
    es_host: str
    es_port: int
    app_host: str
    app_port: int
    min_page_size: int
    max_page_size: int
    max_page_number: int
    max_search_length: int

    class Config:
        env_file = '.env'
        env_file_encoding = 'utf-8'


settings = Settings()

logging_config.dictConfig(LOGGING)
