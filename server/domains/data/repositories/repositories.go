package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/destination"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/participant"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/payment"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/tour"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/tourimage"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/tourreview"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/tourschedule"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/user"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/userdetails"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/village"
)

type Repositories interface {
	GetUser() user.Repository
	GetUserDetails() userdetails.Repository
	GetTour() tour.Repository
	GetTourDestination() destination.Repository
	GetTourSchedule() tourschedule.Repository
	GetTourReview() tourreview.Repository
	GetTourImage() tourimage.Repository
	GetTourParticipation() participant.Repository
	GetTourPayment() payment.Repository
	GetVillage() village.Repository
}

type repo struct {
	user              user.Repository
	userDetails       userdetails.Repository
	tour              tour.Repository
	tourDestination   destination.Repository
	tourSchedule      tourschedule.Repository
	tourReview        tourreview.Repository
	tourImage         tourimage.Repository
	tourParticipation participant.Repository
	tourPayment       payment.Repository
	village           village.Repository
}

// GetVillage implements Repositories.
func (*repo) GetVillage() village.Repository {
	panic("unimplemented")
}

// GetTour implements Repositories.
func (r *repo) GetTour() tour.Repository {
	return r.tour
}

// GetTourDestination implements Repositories.
func (r *repo) GetTourDestination() destination.Repository {
	return r.tourDestination
}

// GetTourImage implements Repositories.
func (r *repo) GetTourImage() tourimage.Repository {
	return r.tourImage
}

// GetTourParticipation implements Repositories.
func (r *repo) GetTourParticipation() participant.Repository {
	return r.tourParticipation
}

// GetTourPayment implements Repositories.
func (r *repo) GetTourPayment() payment.Repository {
	return r.tourPayment
}

// GetTourReview implements Repositories.
func (r *repo) GetTourReview() tourreview.Repository {
	return r.tourReview
}

// GetTourSchedule implements Repositories.
func (r *repo) GetTourSchedule() tourschedule.Repository {
	return r.tourSchedule
}

// GetUser implements Repositories.
func (r *repo) GetUser() user.Repository {
	return r.user
}

// GetUserDetails implements Repositories.
func (r *repo) GetUserDetails() userdetails.Repository {
	return r.userDetails
}

func New(logger *zap.Logger, db *gorm.DB) Repositories {
	rp := new(repo)
	rp.tour = tour.NewGorm(logger, db)
	rp.tourReview = tourreview.NewGorm(logger, db)
	rp.tourDestination = destination.NewGorm(logger, db)
	rp.tourImage = tourimage.NewGorm(logger, db)
	rp.tourSchedule = tourschedule.NewGorm(logger, db)
	rp.tourParticipation = participant.NewGorm(logger, db)
	rp.tourPayment = payment.NewGorm(logger, db)
	rp.user = user.NewGorm(logger, db)
	rp.userDetails = userdetails.NewGorm(logger, db)

	return rp
}
