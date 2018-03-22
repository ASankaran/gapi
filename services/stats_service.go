package services

import (
	"../models"
	"../database"
	"../responses"
	"../cache"
	"strconv"
	"encoding/json"
)

type StatsService struct {
}

func (services StatsService) IncrementStat(category string, attribute string, field string) {
	var stat models.Stat
	stat.Category = category
	stat.Attribute = attribute
	stat.Field = field

	database.DB.Where(map[string]interface{}{"category": category, "attribute": attribute, "field": field}).First(&stat)
	stat.Count += 1
	database.DB.Save(&stat)
}

func (services StatsService) DecrementStat(category string, attribute string, field string) {
	var stat models.Stat
	stat.Category = category
	stat.Attribute = attribute
	stat.Field = field

	database.DB.Where(map[string]interface{}{"category": category, "attribute": attribute, "field": field}).First(&stat)
	stat.Count -= 1
	database.DB.Save(&stat)
}

func (service StatsService) writeBackToCache(key string, stats interface{}) {
	stats_string, _ := json.Marshal(stats)
	cache.WriteString(key, string(stats_string))
}

func (service StatsService) getStats(category string, attribute string) []models.Stat {
	var stats []models.Stat
	database.DB.Where(map[string]interface{}{"category": category, "attribute": attribute}).Find(&stats)
	return stats
}

func (service StatsService) GetRegistrationStats() responses.RegistrationStats {
	var registration_stats responses.RegistrationStats

	if cache.KeyExists("STATS_REGISTRATION") {
		json.Unmarshal([]byte(cache.GetString("STATS_REGISTRATION")), &registration_stats)
		return registration_stats
	}

	registration_stats.School = StatsServices.getStats("REGISTRATION", "school")
	registration_stats.Transportation = StatsServices.getStats("REGISTRATION", "transportation")
	registration_stats.Diet = StatsServices.getStats("REGISTRATION", "diet")
	registration_stats.ShirtSize = StatsServices.getStats("REGISTRATION", "shirtsize")
	registration_stats.Gender = StatsServices.getStats("REGISTRATION", "gender")
	registration_stats.GraduationYear = StatsServices.getStats("REGISTRATION", "graduationyear")
	registration_stats.IsNovice = StatsServices.getStats("REGISTRATION", "isnovice")
	registration_stats.Status = StatsServices.getStats("REGISTRATION", "status")
	registration_stats.Major = StatsServices.getStats("REGISTRATION", "major")
	registration_stats.Attendees = StatsServices.getStats("REGISTRATION", "attendees")

	StatsServices.writeBackToCache("STATS_REGISTRATION", registration_stats)

	return registration_stats
}

func (services StatsService) IncrementRegistrationStats(attendee models.Attendee) {
	StatsServices.IncrementStat("REGISTRATION", "school", attendee.School)
	StatsServices.IncrementStat("REGISTRATION", "transportation", attendee.Transportation)
	StatsServices.IncrementStat("REGISTRATION", "diet", attendee.Diet)
	StatsServices.IncrementStat("REGISTRATION", "shirtsize", attendee.ShirtSize)
	StatsServices.IncrementStat("REGISTRATION", "gender", attendee.Gender)
	StatsServices.IncrementStat("REGISTRATION", "graduationyear", strconv.FormatInt(int64(attendee.GraduationYear), 10))
	if attendee.IsNovice {
		StatsServices.IncrementStat("REGISTRATION", "isnovice", "true")
	} else {
		StatsServices.IncrementStat("REGISTRATION", "isnovice", "false")
	}
	StatsServices.IncrementStat("REGISTRATION", "status", attendee.Status)
	StatsServices.IncrementStat("REGISTRATION", "major", attendee.Major)
	StatsServices.IncrementStat("REGISTRATION", "attendees", "count")

	cache.DeleteString("STATS_REGISTRATION")
}

func (services StatsService) DecrementRegistrationStats(attendee models.Attendee) {
	StatsServices.DecrementStat("REGISTRATION", "school", attendee.School)
	StatsServices.DecrementStat("REGISTRATION", "transportation", attendee.Transportation)
	StatsServices.DecrementStat("REGISTRATION", "diet", attendee.Diet)
	StatsServices.DecrementStat("REGISTRATION", "shirtsize", attendee.ShirtSize)
	StatsServices.DecrementStat("REGISTRATION", "gender", attendee.Gender)
	StatsServices.DecrementStat("REGISTRATION", "graduationyear", strconv.FormatInt(int64(attendee.GraduationYear), 10))
	if attendee.IsNovice {
		StatsServices.DecrementStat("REGISTRATION", "isnovice", "true")
	} else {
		StatsServices.DecrementStat("REGISTRATION", "isnovice", "false")
	}
	StatsServices.DecrementStat("REGISTRATION", "status", attendee.Status)
	StatsServices.DecrementStat("REGISTRATION", "major", attendee.Major)
	StatsServices.DecrementStat("REGISTRATION", "attendees", "count")

	cache.DeleteString("STATS_REGISTRATION")
}
