import time
import re

RUNS = 10000

if __name__ == '__main__':
    with open('data.php') as fh:
        testString = fh.read()

    def do():
        return "576ad4f370014dfb1d0f17b0e6855f22" in testString

    start = time.time()
    for i in range(RUNS):
        _ = do()
    duration = time.time() - start
    print("Python: %.2fs" % duration)
