import re
with open("input") as f:
    file = f.read().splitlines()

##### 1s part

# print(file)
# for i in range(len(file)):
#     file[i] = file[i].split()

# c=0
# for i in range(len(file[0])):
#     op = file[-1][i]
#     res = 1 if op == "*" else 0
#     for j in range(len(file)-1):
#         nb = int(file[j][i])
#         exec("res " + op + "=  nb")
#     c+=res

# print(c)



# ======= Second Part ======= #

width = []
ops = []
ll = file[len(file)-1]

i=0
while i<len(file[0]):
    ops.append(ll[i])
    i+=1
    c=0
    while i<len(file[0]) and ll[i] == ' ':
        i+=1
        c+=1
    if i==len(file[0]):
        c+=1
    width.append(c)

# print(width, ops)

# cols = []
summ=0

offset = 0
for l,w in enumerate(width):
    col = []
    for k in range(w):
        nb=[]
        for i in range(len(file)-1):
            nb.append(file[i][offset+k])
        col.append("".join(nb))
    exec("summ += " + ops[l].join(col))
    offset+=w+1

# print(cols)
print(summ)