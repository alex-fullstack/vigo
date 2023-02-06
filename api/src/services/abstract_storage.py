from abc import abstractmethod, ABC
from typing import Any


class AbstractObjectStorage(ABC):
    @abstractmethod
    async def upload(self, *, key: str, stored_object: Any):
        pass
