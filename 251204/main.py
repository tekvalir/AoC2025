with open("input") as f:
    lines = f.read().splitlines()

for i in range(len(lines)):
    lines[i] = list(lines[i])

def removable(lines, j, i):
    if lines[j][i] != '@':
        return False
    line = lines[j]
    adj=[]
    if i>0:
        adj.append(line[i-1])            
    if i<len(line)-1:
        adj.append(line[i+1])
    if j > 0:
        adj.append(lines[j-1][i])
        if i>0:
            adj.append(lines[j-1][i-1])
        if i<len(line)-1:
            adj.append(lines[j-1][i+1])
    if j < len(lines)-1:
        adj.append(lines[j+1][i])
        if i>0:
            adj.append(lines[j+1][i-1])
        if i<len(line)-1:
            adj.append(lines[j+1][i+1])
    c=0
    for k in adj:
        if k == '@':
            c+=1
    if c<4:
        return True
    return False

# summ=0
def iterate(lines):
    c=0
    for j in range(len(lines)):
        line = lines[j]
        for i in range(len(line)):
            if removable(lines, j, i):
                line[i] = '.'
                c+=1
    return c

c=0
cn=1
while cn>0:
    cn = iterate(lines)
    c+=cn


print(c)

