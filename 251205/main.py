with open("input") as f:
    file = f.read().splitlines()

valid_bot = []
valid_top = []
i=0
while file[i] != "":
    # print(file[i])
    tab = file[i].split("-")
    valid_bot.append(int(tab[0]))
    valid_top.append(int(tab[1]))
    i+=1

c=0
for j in range(i+1,len(file)):
    ing = int(file[j])
    for k,low in enumerate(valid_bot):
        if ing >= low and ing <= valid_top[k]:
            c+=1
            break

print(c)

n_bot = []
n_top = []

def non_overlap(low, high, n_bot, n_top, i):
    if i == len(n_bot):
        return [low], [high]
    if (n_bot[i]>=low and n_bot[i]<=high):
        if n_bot[i] == low:
            if n_top[i]<high:
                nb, nt = non_overlap(n_top[i]+1, high, n_bot, n_top, i+1)
                return nb, nt
            else:
                return [], []
        else:
            if n_top[i]>=high:
                nb, nt = non_overlap(low, n_bot[i]-1, n_bot, n_top, i+1)
                return nb,nt
            else:
                nb1, nt1 = non_overlap(low, n_bot[i]-1, n_bot, n_top, i+1)
                nb2, nt2 = non_overlap(n_top[i]+1, high, n_bot, n_top, i+1)
                return nb1+nb2, nt1+nt2
    if (n_top[i]>=low and n_top[i]<=high):
        if n_top[i] == high:
            return [], []
        else:
            nb, nt = non_overlap(n_top[i]+1, high, n_bot, n_top, i+1)
            return nb, nt
    return non_overlap(low, high, n_bot, n_top, i+1)
        
for i in range(len(valid_bot)):
    low = valid_bot[i]
    high = valid_top[i]
    nb, nt = non_overlap(low, high, n_bot, n_top, 0)
    # print(low, high, n_bot, n_top, len(n_bot))
    print(nb, nt)
    n_bot+=nb
    n_top+=nt

# print(n_bot, n_top)

c=0
for i in range(len(n_bot)):
    low = n_bot[i]
    high = n_top[i]
    c += high-low+1

print(c)

######################## VERIF RANGES
# c=0
# for j in range(len(n_top)):
#     ing = n_bot[i]
#     for k,low in enumerate(valid_bot):
#         if ing >= low and ing <= valid_top[k]:
#             c+=1
#             break

# print(c==len(n_top))
