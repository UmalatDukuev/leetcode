// В жизни каждого приходит время, когда надо отправиться в отпуск.
// Но мы смотрим на календарь и понимаем, что впереди куча встреч, которые не хочется пропускать.
// Необходимо определить день X начала отпуска длиной в V дней так,
// чтобы отгулять весь отпуск в ближайшие P дней и пропустить минимум Y встреч.
// Считаем, что уже завтра первый возможный день отпуска (day: 1)

// # Пример 1
// # ввод
// daysWithMeetings = [{ day: 3, meetings: 1 }, { day: 4, meetings: 3 }, { day: 14, meetings: 3 }, { day: 21, meetings: 3 }, { day: 28, meetings: 1 },] # дни со встречами уже упорядочены
// periodLength = 30 # В какой срок надо отгулять ВЕСЬ отпуск. В данном примере в ближайшие 30 дней
// vacationLength = 7
// # вывод
// [5, 0] # [день X начала отпуска, сколько встреч Y пропустим]

// # Пример 2
// # ввод
// daysWithMeetings = [{ day: 3, meetings: 1 }, { day: 4, meetings: 3 }, { day: 5, meetings: 3 }, { day: 9, meetings: 5 }, { day: 13, meetings: 2 }, { day: 14, meetings: 1 }, { day: 21, meetings: 3 }, { day: 25, meetings: 3 }, { day: 28, meetings: 6 },]
// periodLength = 31
// vacationLength = 14
// # вывод
// [10, 6] # [через сколько дней начало отпуска, сколько встреч пропустим]

package main

import "fmt"

type MeetingDay struct {
	Day      int
	Meetings int
}

var (
	daysWithMeetings = []MeetingDay{
		{Day: 3, Meetings: 1},
		{Day: 4, Meetings: 3},
		{Day: 14, Meetings: 3},
		{Day: 21, Meetings: 3},
		{Day: 28, Meetings: 1},
	}
	periodLength   = 30
	vacationLength = 7
)

func main() {
	missedMeetings := 0
	lastDayToVac := periodLength - vacationLength + 1
	startDay := 1
	lastDay := startDay + vacationLength - 1
	minMissedMeets := 10000000
	resultStartDay := 10000000
	for startDay <= lastDayToVac {
		for _, v := range daysWithMeetings {
			if v.Day >= startDay && v.Day <= lastDay {
				missedMeetings += v.Meetings
			}
		}
		// fmt.Println("-----------------------------------------------")
		// fmt.Println(startDay)
		// fmt.Println(lastDay)
		// fmt.Println(missedMeetings)
		if missedMeetings < minMissedMeets {
			minMissedMeets = missedMeetings
			resultStartDay = startDay
		}
		missedMeetings = 0
		startDay++
		lastDay++
	}
	fmt.Println(resultStartDay)
	fmt.Println(minMissedMeets)
}
