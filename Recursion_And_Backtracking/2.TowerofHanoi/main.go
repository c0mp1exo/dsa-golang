package main

import "fmt"

func main() {
	//disks count
	disks := 3
	TowerOfHanoi(1, 2, 3, disks)

}

//Tower of Hanoi prints instructions to move disk form tower A to tower B using tower C
func TowerOfHanoi(towerA, towerB, towerC, disk int) {
	if disk <= 0 {
		return
	}
	TowerOfHanoi(towerA, towerC, towerB, disk-1)
	fmt.Printf("Moving Disk %d from tower %d to tower %d\n", disk, towerA, towerB)
	TowerOfHanoi(towerC, towerB, towerA, disk-1)
}
