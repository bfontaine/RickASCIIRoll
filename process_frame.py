# -*- coding: UTF-8 -*-

import re
import sys
import argparse

def rewrite_frame(content, lpad=0, rpad=0):
    m = re.findall(r"<pre>(.+)</pre>", content, re.DOTALL)
    if m:
        content = m[0]

    if not rpad and not lpad:
        return content.strip()

    lpadding = "M" * lpad
    rpadding = "M" * rpad

    # Remove black padding
    lines = content.strip().split("\n")
    for i, line in list(enumerate(lines)):
        if line.startswith(lpadding):
            line = line[lpad:]
        if line.endswith(rpadding):
            line = line[:-rpad]

        lines[i] = line

    return "\n".join(lines)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--lpad", type=int, default=0)
    parser.add_argument("--rpad", type=int, default=0)
    opts = parser.parse_args()

    sys.stdout.write(rewrite_frame(sys.stdin.read(), **vars(opts)))
