from typing import Optional, Any

from aioredis import Redis

from services.abstract_storage import AbstractStorage

redis: Optional[Redis] = None


class ElasticSearchRedisStorage(AbstractStorage):
    def __init__(self, *, redis: Redis):
        self.redis = redis

    async def set(self, *, key: str, stored_value: Any):
        await self.redis.set(
            'elastic_cache_{key}'.format(key=key),
            stored_value,
            expire=60,
        )

    async def get(self, *, key: str) -> Any:
        stored_value = await self.redis.get(
            'elastic_cache_{key}'.format(key=key),
        )
        if not stored_value:
            return None
        return stored_value


# Функция понадобится при внедрении зависимостей
async def get_redis() -> Redis:
    return redis
