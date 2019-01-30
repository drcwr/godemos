import time
import re

RUNS = 10000

if __name__ == '__main__':
    # The Complete Works of William Shakespeare by William Shakespeare
    # http://www.gutenberg.org/files/100/100-0.txt
    file = '100-0.txt' # 'data.php'
    with open(file) as fh:
        testString = fh.read()

    def do():
        return "Means to immure herself and not be seen." in testString

    start = time.time()
    for i in range(RUNS):
        _ = do()
    duration = time.time() - start
    print("Python: %.2fs" % duration)
    print(do())
