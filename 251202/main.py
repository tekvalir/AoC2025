with open("input") as f:
    file = f.read().split(",")

ranges = []
for i in file:
    r = i.split("-")
    ranges+=r

summ = 0

print(ranges)

def check_ptrn(nbtxt):
    i = 1
    while i != len(nbtxt)//2+1:
        # print("ptr len ", i)
        if len(nbtxt)%i == 0:
            ptr = nbtxt[:i]
            # print(ptr)
            j=1
            # print("check ", nbtxt[i*j:min(len(nbtxt), i*(j+1))], " == ", ptr, "=> ", nbtxt[i*j:min(len(nbtxt), i*(j+1))]==ptr, min(len(nbtxt), i*(j+1)))
            while j*i<len(nbtxt) and nbtxt[i*j:min(len(nbtxt), i*(j+1))]==ptr:
                j+=1
            # print(j)
            if j*i>=len(nbtxt):
                return True
        i+=1
    return False

for i in range(len(ranges)//2):
    down = int(ranges[2*i])
    up = int(ranges[2*i+1])
    print(down, up)
    for j in range(down, up+1):
        # print("check " + jtxt)
        jtxt = str(j)
        if check_ptrn(jtxt):
            print(jtxt)
            summ+=j
                

print(check_ptrn("11")) 

print(summ)