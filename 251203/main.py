with open("input") as f:
    lines = f.read().splitlines()

summ=0
for line in lines:
    great = [0,0]
    for i in range(len(line)):
        nb = int(line[i])
        if i == len(line)-1 and nb > great[1]:
            great[1] = nb
        elif nb > great[0]:
            great[1] = 0
            great[0]=nb
        elif nb <= great[0] and  nb > great[1]:
            great[1] = nb
    summ += great[0]*10+great[1]
    print(great)
              
print(summ)


def select(n, lines):
    summ=0
    for line in lines:
        battery = [0 for _ in range(n)]
        for i in range(len(line)):
            nb = int(line[i])
            if len(line) - i >= 12:
                j = 0
            else:
                j = 12 - len(line) + i
            for k in range(j,n):
                if nb > battery[k]:
                    battery[k] = nb
                    for l in range(k+1,len(battery)):
                        battery[l] = 0
                    break
        cf = 0
        for i in range(n):
            c = battery[i]
            cf += battery[i]*10**(n-i-1)
        summ += cf
    return summ

print(select(12, lines))