package piscine

func SortWordArr(a []string) {
	for i := len(a) - 1; i > 0; i-- {
		for j := len(a) - 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

/*SortWordArr([09T2o9EwS8rQZ m1n Mn sfORJ7 KQ3neMCY WRWv HvP  gQ1HZrwK t X2Dc3C9 PvP PxMGvQN7aec4q y8qvD AMPr6d  yR 9VEjTqWWc ]) ==
[0 H K P m t y y]
 [09T2o9EwS8rQZ HvP  gQ1HZrwK KQ3neMCY WRWv PxMGvQN7aec4q m1n Mn sfORJ7 t X2Dc3C9 PvP y8qvD AMPr6d  yR 9VEjTqWWc ]*/
