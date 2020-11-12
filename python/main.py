import asyncio
import json
from pathlib import Path
from timeit import default_timer as timer
from typing import NamedTuple

from aiohttp import ClientSession


class Digimon(NamedTuple):
    name: str
    img: str
    level: str


def get_names():
    return Path("digimon.txt").read_text().split("\n")


async def request(url, session):
    response = await session.request(method="GET", url=url)
    parsed, *_ = json.loads(await response.text())
    digimon = Digimon(**parsed)
    print(digimon)


async def main():
    async with ClientSession() as session:
        await asyncio.gather(
            *[
                request(
                    f"https://digimon-api.vercel.app/api/digimon/name/{name}",
                    session
                )
                for name in get_names()
            ]
        )


if __name__ == "__main__":
    start = timer()
    asyncio.run(main())
    end = timer()
    print(f"Duration: {end-start}s")
