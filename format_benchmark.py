import sys
import re


rex = re.compile(r'Benchmark([a-zA-Z0-9]+)_(Jsony|Stdlib)-[0-9]+\s+\d+\s+([0-9.]+)\sns/op')

libs: dict[str, float] = {}
prev = ''


def report() -> None:
    jsony = round(libs["Jsony"])
    stdlib = round(libs["Stdlib"])
    if jsony < stdlib:
        jsony = f'{jsony} 🏆'  # type: ignore
    else:
        stdlib = f'{stdlib} 🏆'  # type: ignore
    print(f'| {prev:12} | {jsony:>8} | {stdlib:>8} | ')


for line in sys.stdin:
    line = line.strip()
    match = rex.fullmatch(line)
    if not match:
        print(line)
        continue
    (cat, lib, dur) = match.groups()
    if cat == prev:
        libs[lib] = float(dur)
        continue
    if prev:
        report()
    else:
        print('| category     | jsony     | stdlib   |')
        print('| ------------ | --------: | -------: |')
    libs = {}
    libs[lib] = float(dur)
    prev = cat

report()
