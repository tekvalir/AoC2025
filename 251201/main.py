with open("input") as f:
    inst = f.read().splitlines()

lock = 50
n0 = 0

for i in inst:
    nb = int(i[1:])
#     print(i[0], nb, lock)
    for _ in range(nb):
        if i[0] == 'L':
            lock-=1
            lock%=100
        elif i[0] == 'R':
            lock+=1
            lock%=100
        else:
            raise Exception()
        if lock == 0:
#             print("ding")
            n0 += 1
    # if lock == 0:
    #     n0 += 1
    #     print("ding!")

# print(lock)
print(n0)
# print(len(inst))