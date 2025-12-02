package payment

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"slices"
)

type CronObject struct{
	Type string
	Value string 
}

type Cronn struct{
	Minute CronObject
	Hour CronObject 
	Day CronObject
	Month CronObject
	Weekday CronObject
} 

func (Cron Cronn) cronDescription() {
	fmt.Printf("Minute %s, type %s \n", Cron.Minute.Value, Cron.Minute.Type )
	fmt.Printf("Hour %s, type %s \n", Cron.Hour.Value, Cron.Hour.Type )
	fmt.Printf("Day %s, type %s \n", Cron.Day.Value, Cron.Day.Type )
	fmt.Printf("Month %s, type %s \n", Cron.Month.Value, Cron.Month.Type )
	fmt.Printf("Weekday %s, type %s \n", Cron.Weekday.Value, Cron.Weekday.Type )

}


func makeRange(min int , max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}


// # TODO 
// # CHECK IF THIS IS GONNA GET IMPLEMENTED! 
func nextWeekDay(Cron Cronn, times []time.Time) time.Duration {
	var day time.Duration
	setTime := times[len(times) - 1]
	switch Cron.Day.Type {
		case "wildcard":
			day = time.Duration(setTime.Day()) + time.Duration(1)
		case "fixed":
			value , err  := strconv.Atoi(Cron.Day.Value)
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			day = time.Duration(value)
		case "range":
			days := strings.Split(Cron.Day.Value, "-")
			min, err := strconv.Atoi(days[0]) 
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}	
			max, err := strconv.Atoi(days[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			// daysRange := makeRange(min, max)	
			if setTime.Day() >= min && setTime.Day() <= max {
				day = time.Duration(setTime.Day()) + time.Duration(1)
			}
			if setTime.Day() + 1 > max || setTime.Day() < min  {
				day = time.Duration(min)
			}		
		case "step":
			days := strings.Split(Cron.Day.Value, "/")
			var from time.Duration
			var intFrom int
			
			step, err := strconv.Atoi(days[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			if days[0] == "*"  {
				from = time.Duration(setTime.Day()) 
			} else 
			{
				intFrom, err = strconv.Atoi(days[0])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				if len(times) > 1 && setTime.Day() > intFrom  {
					from = time.Duration(setTime.Day())
				} else {
					from = time.Duration(intFrom)
				}
			}
			day = from + time.Duration(step)
			// fmt.Printf("from: ---> > > > %d or %d \n", intFrom, int(day))
			if int(day) > 24 && int(day) % 24 >= 1 && int(day) - 24 < intFrom {
				day = time.Duration(intFrom) 
			}
				
		case "list":
			days := strings.Split(Cron.Day.Value, ",")
			var intDays []int
			for minu := range days {
				intMinu, err := strconv.Atoi(days[minu])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				intDays = append(intDays, intMinu)
			}
			sort.Ints(intDays)
			if setTime.Day() >= slices.Max(intDays){
				day = time.Duration(intDays[0])
			} else {
				for i := range intDays {
					// fmt.Printf("from: ---> > > > %d or %d or  %d \n", intDays[i], setTime.Day(), len(times))
					if len(times) < 2 && setTime.Day() == intDays[i] {
						day = time.Duration(intDays[i])
						break
					}
					if setTime.Day() < intDays[i] { 
						day = time.Duration(intDays[i])
						break
					}
				}
			}
	}
	return day
}
   
func nextMonth(Cron Cronn, times []time.Time, firstExecution bool) time.Duration {
	var month time.Duration
	setTime := times[len(times) - 1]
	switch Cron.Month.Type {
		case "wildcard":
			if firstExecution{
				month = time.Duration(int(setTime.Month())) 			
			} else {
				month = time.Duration(int(setTime.Month())) + time.Duration(1)
			}
			
		case "fixed":
			value , err  := strconv.Atoi(Cron.Month.Value)
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			month = time.Duration(value)
		case "range":
			months := strings.Split(Cron.Month.Value, "-")
			min, err := strconv.Atoi(months[0]) 
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}	
			max, err := strconv.Atoi(months[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			// monthsRange := makeRange(min, max)	
			if int(setTime.Month()) >= min && int(setTime.Month()) <= max {
				month = time.Duration(int(setTime.Month())) + time.Duration(1)
			}
			if int(setTime.Month()) + 1 > max || int(setTime.Month()) < min  {
				month = time.Duration(min)
			}		
		case "step":
			months := strings.Split(Cron.Month.Value, "/")
			var from time.Duration
			var intFrom int
			
			step, err := strconv.Atoi(months[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			if months[0] == "*"  {
				from = time.Duration(int(setTime.Month())) 
			} else 
			{
				intFrom, err = strconv.Atoi(months[0])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				if len(times) > 1 && int(setTime.Month()) > intFrom  {
					from = time.Duration(int(setTime.Month()))
				} else {
					from = time.Duration(intFrom)
				}
			}
			month = from + time.Duration(step)
			// fmt.Printf("from: ---> > > > %d or %d \n", intFrom, int(month))
			if int(month) > 24 && int(month) % 24 >= 1 && int(month) - 24 < intFrom {
				month = time.Duration(intFrom) 
			}
				
		case "list":
			months := strings.Split(Cron.Month.Value, ",")
			var intMonths []int
			for minu := range months {
				intMinu, err := strconv.Atoi(months[minu])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				intMonths = append(intMonths, intMinu)
			}
			sort.Ints(intMonths)
			if int(setTime.Month()) >= slices.Max(intMonths){
				month = time.Duration(intMonths[0])
			} else {
				for i := range intMonths {
					// fmt.Printf("from: ---> > > > %d or %d or  %d \n", intMonths[i], int(setTime.Month()), len(times))
					if len(times) < 2 && int(setTime.Month()) == intMonths[i] {
						month = time.Duration(intMonths[i])
						break
					}
					if int(setTime.Month()) < intMonths[i] { 
						month = time.Duration(intMonths[i])
						break
					}
				}
			}
	}
	return month
}


func nextDay(Cron Cronn, times []time.Time, firstExecution bool) time.Duration  {
	var day time.Duration
	setTime := times[len(times) - 1]
	switch Cron.Day.Type {
		case "wildcard":
			if firstExecution {
				day = time.Duration(setTime.Day()) 			
			} else {
				day = time.Duration(setTime.Day()) + time.Duration(1)
			} 
		case "fixed":
			value , err  := strconv.Atoi(Cron.Day.Value)
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			day = time.Duration(value)
		case "range":
			days := strings.Split(Cron.Day.Value, "-")
			min, err := strconv.Atoi(days[0]) 
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}	
			max, err := strconv.Atoi(days[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			// daysRange := makeRange(min, max)	
			if setTime.Day() >= min && setTime.Day() <= max {
				day = time.Duration(setTime.Day()) + time.Duration(1)
			}
			if setTime.Day() + 1 > max || setTime.Day() < min  {
				day = time.Duration(min)
			}		
		case "step":
			days := strings.Split(Cron.Day.Value, "/")
			var from time.Duration
			var intFrom int
			
			step, err := strconv.Atoi(days[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			if days[0] == "*"  {
				from = time.Duration(setTime.Day()) 
			} else 
			{
				intFrom, err = strconv.Atoi(days[0])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				if len(times) > 1 && setTime.Day() > intFrom  {
					from = time.Duration(setTime.Day())
				} else {
					from = time.Duration(intFrom)
				}
			}
			day = from + time.Duration(step)
			// fmt.Printf("from: ---> > > > %d or %d \n", intFrom, int(day))
			if int(day) > 24 && int(day) % 24 >= 1 && int(day) - 24 < intFrom {
				day = time.Duration(intFrom) 
			}
				
		case "list":
			days := strings.Split(Cron.Day.Value, ",")
			var intDays []int
			for minu := range days {
				intMinu, err := strconv.Atoi(days[minu])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				intDays = append(intDays, intMinu)
			}
			sort.Ints(intDays)
			if setTime.Day() >= slices.Max(intDays){
				day = time.Duration(intDays[0])
			} else {
				for i := range intDays {
					// fmt.Printf("from: ---> > > > %d or %d or  %d \n", intDays[i], setTime.Day(), len(times))
					if len(times) < 2 && setTime.Day() == intDays[i] {
						day = time.Duration(intDays[i])
						break
					}
					if setTime.Day() < intDays[i] { 
						day = time.Duration(intDays[i])
						break
					}
				}
			}
	}
	return day
}


func nextHour(Cron Cronn, times []time.Time, firstExecution bool) time.Duration {
	var hour time.Duration
	setTime := times[len(times) - 1]
	switch Cron.Hour.Type {
		case "wildcard":
			if firstExecution {
				hour = time.Duration(setTime.Hour())
			} else {
				hour = time.Duration(setTime.Hour()) + time.Duration(1)
			}
			
		case "fixed":
			value , err  := strconv.Atoi(Cron.Hour.Value)
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			hour = time.Duration(value)
		case "range":
			hours := strings.Split(Cron.Hour.Value, "-")
			min, err := strconv.Atoi(hours[0]) 
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}	
			max, err := strconv.Atoi(hours[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			// hoursRange := makeRange(min, max)	
			if setTime.Hour() >= min && setTime.Hour() <= max {
				hour = time.Duration(setTime.Hour()) + time.Duration(1)
			}
			if setTime.Hour() + 1 > max || setTime.Hour() < min  {
				hour = time.Duration(min)
			}		
		case "step":
			hours := strings.Split(Cron.Hour.Value, "/")
			var from time.Duration
			var intFrom int
			
			step, err := strconv.Atoi(hours[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			if hours[0] == "*"  {
				from = time.Duration(setTime.Hour()) 
			} else 
			{
				intFrom, err = strconv.Atoi(hours[0])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				if len(times) > 1 && setTime.Hour() > intFrom  {
					from = time.Duration(setTime.Hour())
				} else {
					from = time.Duration(intFrom)
				}
			}
			hour = from + time.Duration(step)
			// fmt.Printf("from: ---> > > > %d or %d \n", intFrom, int(hour))
			if int(hour) > 24 && int(hour) % 24 >= 1 && int(hour) - 24 < intFrom {
				hour = time.Duration(intFrom) 
			}
				
		case "list":
			hours := strings.Split(Cron.Hour.Value, ",")
			var intHours []int
			for minu := range hours {
				intMinu, err := strconv.Atoi(hours[minu])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				intHours = append(intHours, intMinu)
			}
			sort.Ints(intHours)
			if setTime.Hour() >= slices.Max(intHours){
				hour = time.Duration(intHours[0])
			} else {
				for i := range intHours {
					// fmt.Printf("from: ---> > > > %d or %d or  %d \n", intHours[i], setTime.Hour(), len(times))
					if len(times) < 2 && setTime.Hour() == intHours[i] {
						hour = time.Duration(intHours[i])
						break
					}
					if setTime.Hour() < intHours[i] { 
						hour = time.Duration(intHours[i])
						break
					}
				}
			}
	}
	return hour
}



func nextMinute(Cron Cronn, times []time.Time, firstExecution bool) time.Duration {
	var minute time.Duration
	setTime := times[len(times) - 1]
	
	switch Cron.Minute.Type {
		case "wildcard":
			if firstExecution {
				minute = time.Duration(setTime.Minute())
			} else {
				minute = time.Duration(setTime.Minute()) + time.Duration(1)
			}
		case "fixed":
			value , err  := strconv.Atoi(Cron.Minute.Value)
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			minute = time.Duration(value)
		case "range":
			minutes := strings.Split(Cron.Minute.Value, "-")
			min, err := strconv.Atoi(minutes[0]) 
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}	
			max, err := strconv.Atoi(minutes[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			// minutesRange := makeRange(min, max)	
			if setTime.Minute() >= min && setTime.Minute() <= max {
				minute = time.Duration(setTime.Minute()) + time.Duration(1)
			}
			if setTime.Minute() + 1 > max || setTime.Minute() < min  {
				minute = time.Duration(min)
			}		
		case "step":
			minutes := strings.Split(Cron.Minute.Value, "/")
			var from time.Duration
			var intFrom int
			
			step, err := strconv.Atoi(minutes[1])
			if err != nil  {
				log.Printf("cannot convert to int %v", err)					
			}
			if minutes[0] == "*"  {
				from = time.Duration(setTime.Minute()) 
			} else 
			{
				intFrom, err = strconv.Atoi(minutes[0])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				if len(times) > 1 && setTime.Minute() > intFrom  {
					from = time.Duration(setTime.Minute())
				} else {
					from = time.Duration(intFrom)
				}
			}
			minute = from + time.Duration(step)
			// fmt.Printf("from: ---> > > > %d or %d \n", intFrom, int(minute))
			if int(minute) > 60 && int(minute) % 60 >= 1 && int(minute) - 60 < intFrom {
				minute = time.Duration(intFrom) 
			}
				
		case "list":
			minutes := strings.Split(Cron.Minute.Value, ",")
			var intMinutes []int
			for minu := range minutes {
				intMinu, err := strconv.Atoi(minutes[minu])
				if err != nil  {
					log.Printf("cannot convert to int %v", err)					
				}
				intMinutes = append(intMinutes, intMinu)
			}
			sort.Ints(intMinutes)
			if setTime.Minute() >= slices.Max(intMinutes){
				minute = time.Duration(intMinutes[0])
			} else {
				for i := range intMinutes {
					// fmt.Printf("from: ---> > > > %d or %d or  %d \n", intMinutes[i], setTime.Minute(), len(times))
					if len(times) < 2 && setTime.Minute() == intMinutes[i] {
						minute = time.Duration(intMinutes[i])
						break
					}
					if setTime.Minute() < intMinutes[i] { 
						minute = time.Duration(intMinutes[i])
						break
					}
				}
			}
	}
	return minute
}



func(Cron Cronn) nextExecutions( n int, times []time.Time) []time.Time {
	var month time.Duration
	var day time.Duration
	var hour time.Duration
	var minute time.Duration
	var firstExecution = false 
	if len(times) == n + 1 {
		return times
	}
	if len(times) == 0 {
		now := time.Now().Local()
		times = append(times, now)
		firstExecution = true
	}
	setTime := times[len(times) - 1]
	if firstExecution {
		minute = nextMinute(Cron, times, true)
		hour = nextHour(Cron, times, true)
		day = nextDay(Cron, times, true)
		month = nextMonth(Cron, times, true)

	}else{
		minute = nextMinute(Cron, times, false)
		hour = time.Duration(setTime.Hour())
		day = time.Duration(setTime.Day())
		month = time.Duration(setTime.Month())
	}
	// fmt.Printf("This is the one being executed ---> %s", setTime.String())
	nextTime := time.Date(setTime.Year(), time.Month(int(month)), int(day), int(hour), int(minute), 0, 0, time.Local)
	if slices.Contains(times, nextTime){
		hour = nextHour(Cron, times, false)
		nextTime = time.Date(setTime.Year(), setTime.Month(), setTime.Day(), int(hour), int(minute), 0, 0, time.Local)
		if slices.Contains(times, nextTime){
			day = nextDay(Cron,times, false)
			nextTime = time.Date(setTime.Year(), setTime.Month(), int(day), int(hour), int(minute), 0, 0, time.Local)
		}
		if slices.Contains(times, nextTime){
			month = nextMonth(Cron, times, false)
			nextTime = time.Date(setTime.Year(), time.Month(int(month)), int(day), int(hour), int(minute), 0, 0, time.Local)
		}
	} 
	times = append(times, nextTime)
	// times = append(times, time.Date(setTime.Year(), setTime.Month(), setTime.Day(), int(hour), int(minute), 0, 0, time.Local))
	return Cron.nextExecutions(n, times) 
}

func cronValidation(cron string)  Cronn {
	var CronValidated Cronn
	var CronObject CronObject
	cronRegex := `^(?P<minute>(?:\*|[0-9]+(?:-[0-9]+)?)(?:/[0-9]+)?(?:,[0-9]+)*) (?P<hour>(?:\*|[0-9]+(?:-[0-9]+)?)(?:/[0-9]+)?(?:,[0-9]+)*) (?P<day>(?:\*|[0-9]+(?:-[0-9]+)?)(?:/[0-9]+)?(?:,[0-9]+)*) (?P<month>(?:\*|[0-9]+(?:-[0-9]+)?)(?:/[0-9]+)?(?:,[0-9]+)*) (?P<weekday>(?:\*|[0-9]+(?:-[0-9]+)?)(?:/[0-9]+)?(?:,[0-9]+)*)$`
	regexes := map[string]*regexp.Regexp{
		"wildcard": regexp.MustCompile(`^\*$`),
		"fixed":    regexp.MustCompile(`^[0-9]+$`),
		"range":    regexp.MustCompile(`^[0-9]+-[0-9]+$`),
		"step":     regexp.MustCompile(`^(?:\*|[0-9]+(?:-[0-9]+)?)/[0-9]+$`),
		"list":     regexp.MustCompile(`^[0-9]+(?:,[0-9]+)+$`),
	}
	re := regexp.MustCompile(cronRegex)
	result := re.FindStringSubmatch(cron)
	if len(result) == 0 {
		panic("cron Validation fails, invalid cron" )
	}
	names := re.SubexpNames()
	for k, v := range names {
		if k != 0 && v != "" {
			value := result[k]
			for typ, re := range regexes {
				if re.MatchString(value) {	
					// fmt.Printf("%s %s %s\n", v,  typ, value)
					CronObject.Type = typ
					CronObject.Value = value
					// fmt.Printf("%s", v)
					switch v {
					case "minute":
						CronValidated.Minute= CronObject
					case "hour":
						CronValidated.Hour= CronObject
					case "day":
						CronValidated.Day= CronObject
					case "month":
						CronValidated.Month= CronObject	
					case "weekday":
						CronValidated.Weekday= CronObject
					}

				}
			}
		}
	} 
	return CronValidated 
}
