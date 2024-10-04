package repositories

import (
    "hotel-booking/models"
    "gorm.io/gorm"
)

type AdminRepository struct {
	dbHandler   *sql.DB
	// transaction *sql.Tx
}

func NewAdminRepository(dbHandler *sql.DB) *AdminRepository {
	return &AdminRepository{
		dbHandler: dbHandler,
	}
}

func (ar AdminsRepository) Authenticate(username, password string) (*models.Admin, *models.ResponseError) {
    var admin models.Admin
    
        // Query to find the admin by username
        query := `SELECT id, username, password, email, created_at FROM admins WHERE username = ?`
        err := ar.dbHandler.QueryRow(query, username).Scan(&admin.ID, &admin.Username, &admin.Password, &admin.Email, &admin.CreatedAt)
    
        if err != nil {
            if err == sql.ErrNoRows {
                return nil, &models.ResponseError{
                    Message: "Invalid username or password",
                    Status:  http.StatusUnauthorized,
                }
            }
            return nil, &models.ResponseError{
                Message: err.Error(),
                Status:  http.StatusInternalServerError,
            }
        }
    
        // Compare the provided password with the hashed password in the database
        err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
        if err != nil {
            return nil, &models.ResponseError{
                Message: "Invalid username or password",
                Status:  http.StatusUnauthorized,
            }
        }
    
        // Optionally update last login in the database
        _, err = ar.dbHandler.Exec(`UPDATE admins SET last_login = ? WHERE id = ?`, time.Now(), admin.ID)
        if err != nil {
            return nil, &models.ResponseError{
                Message: err.Error(),
                Status:  http.StatusInternalServerError,
            }
        }
    
        return &admin, nil
    }
