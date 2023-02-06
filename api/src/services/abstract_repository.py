from abc import abstractmethod, ABC


class AbstractEntityRepository(ABC):
    @abstractmethod
    async def create(self, *, entity: dict) -> str | None:
        pass

    @abstractmethod
    async def get(self, *, entity_id: str) -> dict | None:
        pass

    @abstractmethod
    async def get_list(
        self,
        *,
        page_size: int,
        page_number: int,
        query: dict = None,
        sort: str = '',
    ) -> list[dict]:
        pass
