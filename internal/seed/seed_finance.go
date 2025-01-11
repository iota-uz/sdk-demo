package seed

import (
	"context"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/currency"
	category "github.com/iota-uz/iota-sdk/modules/finance/domain/aggregates/expense_category"
	"github.com/iota-uz/iota-sdk/modules/finance/infrastructure/persistence"
	"github.com/iota-uz/iota-sdk/pkg/application"
)

var Categories = []category.ExpenseCategory{
	category.New(
		"Rent",
		"Office, manufacturing, or other space rented for business purposes.",
		4000,
		&currency.USD,
	),
	category.New(
		"Utilities",
		"Electricity, water, gas, and other utility expenses.",
		1400,
		&currency.USD,
	),
	category.New(
		"Salaries",
		"Monthly payments to employees.",
		25000,
		&currency.USD,
	),
	category.New(
		"Marketing",
		"Costs related to advertising, promotions, and branding.",
		5000,
		&currency.USD,
	),
	category.New(
		"Software Subscriptions",
		"Recurring costs for software licenses and SaaS tools.",
		1200,
		&currency.USD,
	),
	category.New(
		"Office Supplies",
		"Paper, pens, and other general office supply expenses.",
		500,
		&currency.USD,
	),
	category.New(
		"Travel",
		"Business-related travel, including airfare, hotels, and per diems.",
		3000,
		&currency.USD,
	),
	category.New(
		"Insurance",
		"General liability, property, or other insurance policies.",
		2000,
		&currency.USD,
	),
	category.New(
		"Professional Services",
		"Consulting, legal, and accounting services.",
		2500,
		&currency.USD,
	),
	category.New(
		"Training and Development",
		"Workshops, courses, and skill development programs for employees.",
		800,
		&currency.USD,
	),
	category.New(
		"Repairs and Maintenance",
		"Upkeep of office equipment and facilities.",
		600,
		&currency.USD,
	),
	category.New(
		"Taxes",
		"Property taxes, business income taxes, or other tax liabilities.",
		8000,
		&currency.USD,
	),
	category.New(
		"Employee Benefits",
		"Healthcare, retirement plans, and other employee perks.",
		3500,
		&currency.USD,
	),
	category.New(
		"Internet and Communications",
		"Internet, phone, and other communication services.",
		700,
		&currency.USD,
	),
	category.New(
		"Equipment Lease",
		"Lease payments for computers, machinery, or vehicles.",
		1500,
		&currency.USD,
	),
	category.New(
		"Raw Materials",
		"Cost of materials required for production.",
		10000,
		&currency.USD,
	),
	category.New(
		"Shipping and Delivery",
		"Postage, freight, and logistics expenses.",
		2000,
		&currency.USD,
	),
	category.New(
		"R&D",
		"Research and development for new products or services.",
		5000,
		&currency.USD,
	),
	category.New(
		"IT Support",
		"IT infrastructure maintenance and technical support services.",
		1800,
		&currency.USD,
	),
	category.New(
		"Miscellaneous Expenses",
		"Uncategorized small business expenses.",
		300,
		&currency.USD,
	),
}

func createOrUpdateCategory(ctx context.Context, cat category.ExpenseCategory) error {
	categoryRepository := persistence.NewExpenseCategoryRepository()
	matches, err := categoryRepository.GetPaginated(ctx, &category.FindParams{
		Field: "name",
		Query: cat.Name(),
	})
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		_, err = categoryRepository.Create(ctx, cat)
	} else {
		_, err = categoryRepository.Update(ctx, category.NewWithID(
			matches[0].ID(),
			cat.Name(),
			cat.Description(),
			cat.Amount(),
			cat.Currency(),
			cat.CreatedAt(),
			cat.UpdatedAt(),
		))
	}
	return err
}

func CreateExpenseCategories(ctx context.Context, app application.Application) error {
	for _, cat := range Categories {
		if err := createOrUpdateCategory(ctx, cat); err != nil {
			return err
		}
	}
	return nil
}
