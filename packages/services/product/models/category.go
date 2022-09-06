package models

import "time"

type Category interface {
	Details() CategoryDetails
	IsArchived() bool
}

type CategoryDetails struct {
	Id              int                  `json:"id" gorm:"primaryKey"`
	ActiveEndDate   time.Time            `json:"active_end_date"`
	ActiveStartDate time.Time            `json:"active_start_date"`
	Attributes      map[string]Attribute `json:"attributes"`
	Title           string               `json:"title"`
	Description     string               `json:"description"`
	DisplayTemplate string               `json:"display_templates"`
	ExternalId      string               `json:"external_id"`
	Slug            string               `json:"slug"`
	URL             string               `json:"url"`
	MetaTitle       string               `json:"meta_title"`
	MetaDescription string               `json:"meta_description"`
}

type CategoryAsset struct {
	Id      int       `json:"id" gorm:"primaryKey"`
	Sorting string    `json:"sorting"`
	Type    AssetType `json:"type"`
	URL     string    `json:"url"`
	AltText string    `json:"alt_text"`
	Title   string    `json:"title"`
}
