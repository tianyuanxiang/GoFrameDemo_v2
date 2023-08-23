// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo/internal/dao/internal"
)

// internalUserinformationDao is internal type for wrapping internal DAO implements.
type internalUserinformationDao = *internal.UserinformationDao

// userinformationDao is the data access object for table userinformation.
// You can define custom methods on it to extend its functionality as you wish.
type userinformationDao struct {
	internalUserinformationDao
}

var (
	// Userinformation is globally public accessible object for table userinformation operations.
	Userinformation = userinformationDao{
		internal.NewUserinformationDao(),
	}
)

// Fill with you ideas below.
