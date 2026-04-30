package main

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" && a.State == "" && a.Zip == "" {
		return "No address provided"
	}
	return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
}

type ContactInfo struct {
	PhoneNumber string
	Email       string
}

// Attaching a method DisplayContact() to the ContactInfo struct that returns the contact information
// as a formatted string.
func (ci ContactInfo) DisplayContact() string {
	if ci.PhoneNumber == "" && ci.Email == "" {
		return "No contact information provided"
	}
	return "Phone: " + ci.PhoneNumber + ", Email: " + ci.Email
}

// Company struct embeds Address and ContactInfo structs, allowing access
// to their fields and methods directly.
type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetCompanyInfo() string {
	return "Company Name: " + c.Name + "\n" +
		"Business Type: " + c.BusinessType + "\n" +
		"Address: " + c.FullAddress() + "\n" + // Direct access to method of embedded Address struct
		"Email: " + c.Email + "\n" + // Direct access to field of embedded ContactInfo struct
		"Contact Info: " + c.DisplayContact()
}

type CompanyWithOwnEmail struct {
	Company
	Email string // This field will shadow the Email field from ContactInfo
}

func main() {
	company := Company{
		Name: "Tech Innovators Inc.",
		Address: Address{
			Street: "789 Innovation Way",
			City:   "Techville",
			State:  "CA",
			Zip:    "54321",
		},
		ContactInfo: ContactInfo{
			PhoneNumber: "555-1234",
			Email:       "info@techinnovators.com",
		},
		BusinessType: "Software Development",
	}

	println("\n--- Company Info ---")
	println(company.GetCompanyInfo())

	companyWithOwnEmail := CompanyWithOwnEmail{
		Company: company,
		Email:   "contact@techinnovators.com", // This will shadow the Email field from ContactInfo
	}

	println("\n--- Company with Own Email ---")
	println(companyWithOwnEmail.GetCompanyInfo())
	println("Company's own email: " + companyWithOwnEmail.Email)                   // Accessing the shadowed Email field
	println("ContactInfo email: " + companyWithOwnEmail.Company.ContactInfo.Email) // Accessing the original Email field from ContactInfo
}
