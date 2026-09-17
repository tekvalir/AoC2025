with open("input") as f:
    lines = f.read().splitlines()

lines = list(map(list, lines))
print(lines)

hist = {}

def rec_beam(lines, col, it):
    if it == 0:
        for i in range(len(lines[0])):
            if lines[0][i] == 'S':
                print(f"Start at %d, passed to line 1" % i)
                return rec_beam(lines, i, 1)
    if lines[it-1][col] == "|":
        return 0
    lines[it-1][col] = '|'
    print(lines[it-1])
    print(lines[it])
    if it == len(lines)-1:
        print("Bottom reached, return 0")
        return 0
    if lines[it][col] == '^':
        c=1
        if col > 0:
            print(f"Split left, it %d, next col %d" % (it, col-1))
            c+=rec_beam(lines, col-1, it+1)
        if col < len(lines[it])-1:
            print(f"Split right, it %d, next col %d" % (it, col+1))
            c+=rec_beam(lines, col+1, it+1)
        return c
    print(f"No splitter found, it %d, col %d" % (it, col))
    return rec_beam(lines, col, it+1)

print(rec_beam(lines, 0, 0))


def rec_beam2(lines, col, it):
    if it == 0:
        for i in range(len(lines[0])):
            if lines[0][i] == 'S':
                print(f"Start at %d, passed to line 1" % i)
                return rec_beam2(lines, i, 1)
    if it == len(lines)-1:
        print("Bottom reached, return 0")
        return 1
    if lines[it][col] == '^':
        c=0
        if col > 0:
            print(f"Split left, it %d, next col %d" % (it, col-1))
            c+=hist_or_check(lines, col-1, it+1)
        if col < len(lines[it])-1:
            print(f"Split right, it %d, next col %d" % (it, col+1))
            c+=hist_or_check(lines, col+1, it+1)
        return c
    print(f"No splitter found, it %d, col %d" % (it, col))
    return hist_or_check(lines, col, it+1)

def hist_or_check(lines, col, it):
    if it not in hist:
        hist[it] = {}
    if col in hist[it]:
        print("Found in history")
        return hist[it][col]
    print("Not found in history")
    comp = rec_beam2(lines, col, it)
    hist[it][col] = comp
    return comp

# print(rec_beam2(lines, 0, 0))
# print(hist)