package rules

type Character struct {
	Id                 int
	Name               string
	HitPoints          int
	Speed              int
	AC                 int
	TouchAC            int
	FlatFootAC         int
	Classes            string
	Strength           int
	Dexterity          int
	Constitution       int
	Intelligence       int
	Wisdom             int
	Charisma           int
	InitiativeMod      int
	FortitudeSave      int
	ReflexSave         int
	WillSave           int
	AttackHandheld     int
	AttackMissile      int
	GrappleCheck       int
	CarryingLight      int
	CarryingMedium     int
	CarryingHeavy      int
	CarryLiftOverhead  int
	CarryLiftOffGrouhd int
	CarryPushDrag      int
	Languages          string
	Feats              string
	Skills             string
	Background         string
	Description        string
	MostRecent         bool
}
