with open("input") as f:
    lines = f.read().splitlines()

def distance(c1, c2):
    dx = c1[0]-c2[0]
    dy = c1[1]-c2[1]
    dz = c1[2]-c2[2]
    return dx*dx+dy*dy+dz*dz

for i in range(len(lines)):
    lines[i] = lines[i].split(",")
    lines[i] = list(map(int, lines[i]))

distances = []
for i in range(len(lines)):
    distances.append([0 for _ in range(len(lines))])

for i in range(len(distances)):
    for j in range(i+1, len(distances)):
        distances[i][j] = distance(lines[i], lines[j])

junctions = {}
circuits = [0 for _ in range(len(distances))]
nextid=1
n = 1000
c=0
while c<n:
    min_val = -1
    couple = (0,0)
    for i in range(len(distances)):
        for j in range(i+1, len(distances)):
            if i not in junctions:
                junctions[i] = []
            if j in junctions[i]:
                continue
            if min_val == -1 or distances[i][j] < min_val:
                min_val = distances[i][j]
                couple = (i,j)
    i,j = couple
    junctions[i].append(j)
    if circuits[i] == 0 and circuits[j] == 0:
        circuits[i] = nextid
        circuits[j] = nextid
        nextid+=1
    elif circuits[i] != 0 and circuits[j] == 0:
        circuits[j] = circuits[i]
    elif circuits[i] == 0 and circuits[j] != 0:
        circuits[i] = circuits[j]
    else:
        exCircuit = circuits[j]
        for k in range(len(circuits)):
            if circuits[k] == exCircuit :
                circuits[k] = circuits[i]
    c+=1

print(distances)

print(circuits)
circLen = [0 for _ in range(nextid)]

for i in range(len(circuits)):
    circLen[circuits[i]]+=1

trouple = [0,0,0]
for i in circLen[1:]:
    if i > trouple[0]:
        trouple[2] = trouple[1]
        trouple[1] = trouple[0]
        trouple[0] = i
    elif i > trouple[1]:
        trouple[2] = trouple[1]
        trouple[1] = i
    elif i > trouple[2]:
        trouple[2] = i

print(circLen)
print(trouple)
print(trouple[0]*trouple[1]*trouple[2])