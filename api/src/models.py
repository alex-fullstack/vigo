from uuid import UUID

import orjson
from pydantic import BaseModel

from spec import SpecificationEBMLHeader, SpecificationSegment


def orjson_dumps(dump_data, *, default):
    return orjson.dumps(dump_data, default=default).decode()


class ORJsonMixin(BaseModel):
    class Config:
        json_loads = orjson.loads
        json_dumps = orjson_dumps


class SpecificationElement(BaseModel):
    position: int


class Specification(BaseModel):
    ebml_header:  list[SpecificationEBMLHeader]
    segment: list[SpecificationSegment]


class File:
    specification: Specification



