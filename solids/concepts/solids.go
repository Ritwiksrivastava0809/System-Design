package solids

import (
	"context"
	"errors"
	"fmt"
)

// SOLID Principles
//
// S — Single Responsibility Principle (SRP)
// A type, function, or module should have only one reason to change.
//
// O — Open/Closed Principle (OCP)
// Software entities should be open for extension
// but closed for modification.
//
// L — Liskov Substitution Principle (LSP)
// Implementations should be replaceable through their abstractions
// without breaking expected behavior.
//
// I — Interface Segregation Principle (ISP)
// Clients should not be forced to depend on methods they do not use.
//
// D — Dependency Inversion Principle (DIP)
// High-level modules should depend on abstractions,
// not concrete implementations.

// ----------------------------------------------------------------------
// Domain Model
// ----------------------------------------------------------------------

type User struct {
	ID    int64
	Name  string
	Email string
}

// ----------------------------------------------------------------------
// Repository Abstraction
// ----------------------------------------------------------------------

// UserRepository defines the behavior required
// for persisting users.
//
// The service layer depends on this abstraction
// instead of a concrete database implementation.
//
// This follows the Dependency Inversion Principle.
type UserRepository interface {
	Create(ctx context.Context, user User) error
}

// ----------------------------------------------------------------------
// Service Layer
// ----------------------------------------------------------------------

// UserService is a high-level business module.
//
// It depends on the UserRepository abstraction
// rather than a concrete database implementation.
//
// This makes the service:
// - easier to test
// - loosely coupled
// - flexible to future changes
type UserService struct {
	repo UserRepository
}

// NewUserService injects the repository dependency.
//
// This is constructor injection,
// which is idiomatic in Go.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser handles the business workflow
// for creating a user.
//
// The service delegates persistence responsibility
// to the repository layer.
func (s *UserService) CreateUser(
	ctx context.Context,
	user User,
) error {

	// business logic can be added here
	// example:
	// - validation
	// - password hashing
	// - audit logging
	// - event publishing

	return s.repo.Create(ctx, user)
}

// ----------------------------------------------------------------------
// Concrete Repository Implementation
// ----------------------------------------------------------------------

// PostgresUserRepository is a concrete implementation
// of UserRepository.
//
// It contains database-specific logic.
type PostgresUserRepository struct {
	// db *sql.DB
	// other database connection fields
}

// Create persists the user into PostgreSQL.
//
// This implementation detail is hidden from the service layer.
func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user User,
) error {

	// database insert logic here

	return nil
}

type MongoUserRepository struct {
	// db *mongo.Client
	// other database connection fields
}

func (r *MongoUserRepository) Create(
	ctx context.Context,
	user User,
) error {
	// database insert logic here

	return nil
}

// ----------------------------------------------------------------------
// Why This Design Is Better
// ----------------------------------------------------------------------

/*
BAD DESIGN (tight coupling)

func Create(ctx context.Context, user *User) error {
	return db.Create(ctx, user)
}

Problem:
- business logic directly depends on database implementation
- difficult to test
- tightly coupled
- hard to replace database/storage layer

--------------------------------------------------------

GOOD DESIGN (dependency inversion)

UserService ---> UserRepository(interface) ---> PostgreSQL

Benefits:
- loose coupling
- easier testing
- flexible architecture
- clean separation of concerns
- easier maintenance

The service depends only on behavior:

    Create(ctx, user)

It does not care:
- which database is used
- how persistence works
- whether data is stored in PostgreSQL, MongoDB, or memory

This is the core idea behind
the Dependency Inversion Principle.
*/

//===================================================================OCP===================================================================

// Open/Closed Principle (OCP)
//
// Software entities (classes, modules, functions, etc.)
// should be open for extension but closed for modification.
//
// This means you should be able to add new functionality
// without changing existing code.
//
// In Go, we can achieve this through interfaces and composition.
//
// Example:
// - Define an interface for a shape
// - Implement different shapes (Circle, Rectangle) that satisfy the interface
// - Add new shapes without modifying existing code

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// To add a new shape, we simply implement the Shape interface
// without modifying existing code.
// This adheres to the Open/Closed Principle.
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// The function to calculate total area can work with any shape
// without modification, adhering to OCP.
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// The main function demonstrates how we can use the TotalArea function
// with different shapes without changing the TotalArea implementation.
// This design allows us to extend functionality (adding new shapes)
// without modifying existing code, following the Open/Closed Principle.
func ShapeMain() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 3, Height: 4},
	}

	fmt.Printf("Total Area: %.2f\n", TotalArea(shapes))
}

// ===================================================================
// Open/Closed Principle (OCP)
// ===================================================================
//
// Software entities should be:
// - open for extension
// - closed for modification
//
// Meaning:
//
// We should be able to add new functionality
// without modifying stable existing code.
//
// In Go, OCP is commonly achieved using:
// - interfaces
// - composition
// - dependency injection
//
// ===================================================================
// BAD DESIGN — Violates OCP
// ===================================================================

type PaymentService struct{}

// Problem:
//
// Every time a new provider is added:
// - Stripe
// - PayPal
// - Razorpay
// - ApplePay
//
// we must MODIFY this function.
//
// This creates:
// - growing if-else chains
// - regression risk
// - merge conflicts
// - poor maintainability
func (s *PaymentService) Pay(
	amount float64,
	provider string,
) error {

	if provider == "stripe" {

		fmt.Println("Processing payment with Stripe")

	} else if provider == "paypal" {

		fmt.Println("Processing payment with PayPal")

	} else {

		return fmt.Errorf(
			"unsupported payment provider: %s",
			provider,
		)
	}

	return nil
}

// ===================================================================
// GOOD DESIGN — Follows OCP
// ===================================================================

// PaymentProvider defines the behavior required
// for any payment provider.
//
// Any new provider only needs to implement this interface.
type PaymentProvider interface {
	Pay(amount float64) error
}

// ===================================================================
// Stripe Implementation
// ===================================================================

type StripeProvider struct{}

func (s *StripeProvider) Pay(amount float64) error {

	fmt.Printf(
		"Processing ₹%.2f payment using Stripe\n",
		amount,
	)

	return nil
}

// ===================================================================
// PayPal Implementation
// ===================================================================

type PayPalProvider struct{}

func (p *PayPalProvider) Pay(amount float64) error {

	fmt.Printf(
		"Processing ₹%.2f payment using PayPal\n",
		amount,
	)

	return nil
}

// ===================================================================
// Razorpay Implementation
// ===================================================================

// Notice:
//
// We are EXTENDING the system
// without modifying existing service logic.
//
// This is the essence of OCP.
type RazorpayProvider struct{}

func (r *RazorpayProvider) Pay(amount float64) error {

	fmt.Printf(
		"Processing ₹%.2f payment using Razorpay\n",
		amount,
	)

	return nil
}

// ===================================================================
// Payment Service
// ===================================================================

// PaymentServiceV2 depends on abstraction
// instead of concrete implementations.
//
// This follows:
// - OCP
// - Dependency Inversion Principle (DIP)
type PaymentServiceV2 struct {
	provider PaymentProvider
}

// Constructor Injection
//
// The payment provider is injected from outside.
//
// This allows us to switch providers
// without modifying service logic.
func NewPaymentServiceV2(
	provider PaymentProvider,
) *PaymentServiceV2 {

	return &PaymentServiceV2{
		provider: provider,
	}
}

// Checkout contains business workflow.
//
// The service delegates payment processing
// to whichever provider implementation
// was injected.
func (s *PaymentServiceV2) Checkout(
	amount float64,
) error {

	return s.provider.Pay(amount)
}

// ===================================================================
// Main Function
// ===================================================================

func Paymain() {

	// ---------------------------------------------------------------
	// Stripe Example
	// ---------------------------------------------------------------

	stripeProvider := &StripeProvider{}

	stripeService := NewPaymentServiceV2(
		stripeProvider,
	)

	stripeService.Checkout(1500)

	// ---------------------------------------------------------------
	// PayPal Example
	// ---------------------------------------------------------------

	paypalProvider := &PayPalProvider{}

	paypalService := NewPaymentServiceV2(
		paypalProvider,
	)

	paypalService.Checkout(2200)

	// ---------------------------------------------------------------
	// Razorpay Example
	// ---------------------------------------------------------------

	razorpayProvider := &RazorpayProvider{}

	razorpayService := NewPaymentServiceV2(
		razorpayProvider,
	)

	razorpayService.Checkout(5000)
}

// ===================================================================
// Why This Design Is Better
// ===================================================================

/*
Benefits of OCP-Compliant Design:

1. Extensibility
	New payment providers can be added
	without changing existing business logic.

2. Maintainability
	Stable code remains untouched.

3. Lower Regression Risk
	Existing payment flow is less likely to break.

4. Easier Testing
	Mock providers can be injected during tests.

5. Better Separation of Concerns
	Each provider owns its own payment logic.

6. Cleaner Architecture
	Business workflow remains independent
	from infrastructure/provider details.

------------------------------------------------------------

Most Important Insight:

BAD DESIGN:
    Modify existing logic for every new provider

GOOD DESIGN:
    Extend system by adding new implementations

That is the essence of
the Open/Closed Principle.
*/

// ===================================================================
// BAD DESIGN — Violates OCP
// ===================================================================

// Problem:
//
// Every new notification channel requires modifying
// SendNotification().
//
// This creates:
// - growing if-else chains
// - regression risk
// - harder maintenance
// - tightly coupled logic
type Notifier struct{}

func (n *Notifier) SendNotification(
	message string,
	channel string,
) error {

	if channel == "email" {

		fmt.Printf(
			"Sending EMAIL notification: %s\n",
			message,
		)

	} else if channel == "sms" {

		fmt.Printf(
			"Sending SMS notification: %s\n",
			message,
		)

	} else {

		return fmt.Errorf(
			"unsupported notification channel: %s",
			channel,
		)
	}

	return nil
}

// ===================================================================
// GOOD DESIGN — Follows OCP
// ===================================================================

// NotificationChannel defines behavior
// required for sending notifications.
//
// Any new notification channel only needs
// to implement this interface.
type NotificationChannel interface {
	Send(message string) error
}

// ===================================================================
// Email Notification Implementation
// ===================================================================

type EmailNotifier struct{}

func (e *EmailNotifier) Send(
	message string,
) error {

	fmt.Printf(
		"Sending EMAIL notification: %s\n",
		message,
	)

	return nil
}

// ===================================================================
// SMS Notification Implementation
// ===================================================================

type SMSNotifier struct{}

func (s *SMSNotifier) Send(
	message string,
) error {

	fmt.Printf(
		"Sending SMS notification: %s\n",
		message,
	)

	return nil
}

// ===================================================================
// Push Notification Implementation
// ===================================================================

// New notification channels can be added
// WITHOUT modifying existing service logic.
//
// This is the essence of OCP.
type PushNotifier struct{}

func (p *PushNotifier) Send(
	message string,
) error {

	fmt.Printf(
		"Sending PUSH notification: %s\n",
		message,
	)

	return nil
}

// ===================================================================
// Notification Service
// ===================================================================

// NotificationService depends on abstraction
// instead of concrete implementations.
//
// This follows:
// - Open/Closed Principle
// - Dependency Inversion Principle
type NotificationService struct {
	channel NotificationChannel
}

// Constructor Injection
//
// The notification channel is injected externally.
func NewNotificationService(
	channel NotificationChannel,
) *NotificationService {

	return &NotificationService{
		channel: channel,
	}
}

// Notify handles notification workflow.
//
// The actual sending logic is delegated
// to the injected channel implementation.
func (s *NotificationService) Notify(
	message string,
) error {

	return s.channel.Send(message)
}

// ===================================================================
// Main Function
// ===================================================================

func NotificationMain() {

	// ---------------------------------------------------------------
	// Email Notification
	// ---------------------------------------------------------------

	emailService := NewNotificationService(
		&EmailNotifier{},
	)

	emailService.Notify(
		"Hello via Email!",
	)

	// ---------------------------------------------------------------
	// SMS Notification
	// ---------------------------------------------------------------

	smsService := NewNotificationService(
		&SMSNotifier{},
	)

	smsService.Notify(
		"Hello via SMS!",
	)

	// ---------------------------------------------------------------
	// Push Notification
	// ---------------------------------------------------------------

	pushService := NewNotificationService(
		&PushNotifier{},
	)

	pushService.Notify(
		"Hello via Push Notification!",
	)
}

// ===================================================================
// Why This Design Is Better
// ===================================================================

/*
Benefits:

1. Open for Extension
	We can add:
	- WhatsAppNotifier
	- SlackNotifier
	- DiscordNotifier
	- WebhookNotifier

    without changing existing code.

2. Closed for Modification
    NotificationService remains stable.

3. Easier Testing
    Mock channels can be injected.

4. Better Maintainability
    Each notifier owns its own logic.

5. Cleaner Architecture
    Business workflow is separated
    from delivery mechanism.

------------------------------------------------------------

Most Important Insight:

BAD DESIGN:
    Existing logic changes for every new feature

GOOD DESIGN:
    New behavior added through new implementations

That is the core idea behind
the Open/Closed Principle.
*/

//==================================================================LSP==================================================================

// ===================================================================
// Liskov Substitution Principle (LSP)
// ===================================================================
//
// LSP states:
//
// Objects of derived types should be replaceable
// with objects of their base type
// WITHOUT breaking expected behavior.
//
// In Go:
//
// Any implementation of an interface should behave
// in a way that callers expect.
//
// LSP is NOT just about matching method signatures.
//
// It is about preserving:
// - behavioral correctness
// - expected contracts
// - predictable system behavior

// ===================================================================
// Authentication Interface
// ===================================================================

// AuthenticationService defines expected behavior
// for authenticating users.
//
// Any implementation must:
// - return true for valid credentials
// - return false for invalid credentials
// - return errors only for system failures
//
// Implementations should NOT:
// - panic unexpectedly
// - return inconsistent behavior
type AuthenticationService interface {
	Authenticate(
		username string,
		password string,
	) (bool, error)
}

// ===================================================================
// Database Authentication Implementation
// ===================================================================

type DatabaseAuthService struct{}

// Authenticate validates credentials
// using a database-backed authentication system.
func (s *DatabaseAuthService) Authenticate(
	username string,
	password string,
) (bool, error) {

	if username == "" || password == "" {
		return false, errors.New(
			"username and password required",
		)
	}

	// simulate successful authentication
	if username == "admin" &&
		password == "password123" {

		return true, nil
	}

	return false, nil
}

// ===================================================================
// API Authentication Implementation
// ===================================================================

type APIAuthService struct{}

// Authenticate validates credentials
// using an external authentication API.
//
// IMPORTANT:
//
// This implementation preserves the SAME behavioral contract
// as DatabaseAuthService.
//
// Therefore it satisfies LSP.
func (s *APIAuthService) Authenticate(
	username string,
	password string,
) (bool, error) {

	if username == "" || password == "" {
		return false, errors.New(
			"username and password required",
		)
	}

	// simulate successful authentication
	if username == "admin" &&
		password == "password123" {

		return true, nil
	}

	return false, nil
}

// ===================================================================
// BAD IMPLEMENTATION — Violates LSP
// ===================================================================

// BrokenAuthService technically satisfies interface
// BUT breaks expected behavior.
type BrokenAuthService struct{}

func (s *BrokenAuthService) Authenticate(
	username string,
	password string,
) (bool, error) {

	// BAD:
	// panic violates expected contract
	panic("authentication system crashed")
}

// ===================================================================
// Login Workflow
// ===================================================================

// LoginUser works with ANY AuthenticationService.
//
// As long as implementations preserve expected behavior,
// this function works correctly.
//
// This is LSP in action.
func LoginUser(
	auth AuthenticationService,
	username string,
	password string,
) {

	ok, err := auth.Authenticate(
		username,
		password,
	)

	if err != nil {
		fmt.Println(
			"authentication error:",
			err,
		)
		return
	}

	if !ok {
		fmt.Println("invalid credentials")
		return
	}

	fmt.Println("login successful")
}

// ===================================================================
// Main Function
// ===================================================================

func LSPMain() {

	// ---------------------------------------------------------------
	// Database Authentication
	// ---------------------------------------------------------------

	dbAuth := &DatabaseAuthService{}

	LoginUser(
		dbAuth,
		"admin",
		"password123",
	)

	// ---------------------------------------------------------------
	// API Authentication
	// ---------------------------------------------------------------

	apiAuth := &APIAuthService{}

	LoginUser(
		apiAuth,
		"admin",
		"password123",
	)

	// ---------------------------------------------------------------
	// Broken Implementation
	// ---------------------------------------------------------------

	// This compiles...
	// BUT violates LSP behavior contract.

	// brokenAuth := &BrokenAuthService{}
	// LoginUser(
	// 	brokenAuth,
	// 	"admin",
	// 	"password123",
	// )
}

// ===================================================================
// Why This Design Matters
// ===================================================================

/*
GOOD LSP DESIGN:

Code depending on AuthenticationService
can safely use ANY implementation
without unexpected behavior changes.

Benefits:
- predictable behavior
- safer abstraction
- interchangeable implementations
- easier testing
- cleaner architecture

------------------------------------------------------------

IMPORTANT INSIGHT:

LSP is NOT:

    "Does it compile?"

LSP IS:

    "Does it preserve expected behavior?"

------------------------------------------------------------

BAD IMPLEMENTATION:

Compiles successfully
BUT breaks behavioral expectations.

That is an LSP violation.

------------------------------------------------------------

Most Important Takeaway:

Interfaces define method signatures.

LSP defines behavioral correctness.
*/

//==================================================================ISP==================================================================

// ===================================================================
// Interface Segregation Principle (ISP)
// ===================================================================
//
// ISP states:
//
// Clients should not be forced to depend on methods
// they do not use.
//
// In Go, this means:
// - interfaces should be small and focused
// - clients should only depend on the methods they need
//
// Benefits of ISP:
// - better modularity
// - easier testing
// - clearer contracts
// - reduced coupling
// - improved maintainability

// ===================================================================
// Monolithic Interface — Violates ISP
// ===================================================================

// This interface is too broad.
//
// Clients that only need to send emails
// are forced to implement SMS and Push methods.
//
// This creates:
// - unnecessary dependencies
// - harder testing
// - more complex implementations
type NotifierV1 interface {
	SendEmail(message string) error
	SendSMS(message string) error
	SendPush(message string) error
}

// ===================================================================
// Segregated Interfaces — Follows ISP
// ===================================================================

// Each interface is focused on a specific notification channel.
//
// Clients can implement only the interfaces they need,
// adhering to the Interface Segregation Principle.
type EmailNotifierV2 interface {
	SendEmail(message string) error
}

type SMSNotifierV2 interface {
	SendSMS(message string) error
}

type PushNotifierV2 interface {
	SendPush(message string) error
}

// ===================================================================
// Implementations
// ===================================================================

type EmailServiceV1 struct{}

func (e *EmailServiceV1) SendEmail(
	message string,
) error {

	fmt.Printf(
		"Sending EMAIL notification: %s\n",
		message,
	)

	return nil
}

type SMSServiceV1 struct{}

func (s *SMSServiceV1) SendSMS(
	message string,
) error {

	fmt.Printf(
		"Sending SMS notification: %s\n",
		message,
	)

	return nil
}

type PushServiceV1 struct{}

func (p *PushServiceV1) SendPush(
	message string,
) error {

	fmt.Printf(
		"Sending PUSH notification: %s\n",
		message,
	)

	return nil
}

// ===================================================================
// Main Function
// ===================================================================

func ISPMain() {

	email := &EmailService{}
	sms := &SMSService{}
	push := &PushService{}

	email.SendEmail("Hello via Email!")
	sms.SendSMS("Hello via SMS!")
	push.SendPush("Hello via Push Notification!")
}

// ===================================================================
// Why This Design Matters
// ===================================================================

/*
Benefits of ISP-Compliant Design:

1. Focused Interfaces
	Clients only depend on what they need.

2. Easier Testing
	Mock implementations can be created for specific interfaces.

3. Clearer Contracts
	Each interface has a single responsibility.

4. Reduced Coupling
	Changes to one interface do not affect clients of another.

5. Improved Maintainability
	Smaller interfaces are easier to understand and maintain.

------------------------------------------------------------

Most Important Insight:

ISP is about creating focused, cohesive interfaces that clients can depend on without being forced to implement methods they do not use.
*/

// ================================================================
// Interface Segregation Principle (ISP)
// ================================================================
//
// ISP:
// Clients should not depend on methods they do not use.
//
// In Go:
// - keep interfaces small
// - define interfaces around behavior
// - prefer focused contracts
//
// Go proverb:
// "The bigger the interface, the weaker the abstraction."
//

// ================================================================
// Focused Interfaces
// ================================================================

// Email capability
type EmailSender interface {
	SendEmail(message string) error
}

// SMS capability
type SMSSender interface {
	SendSMS(message string) error
}

// Push notification capability
type PushSender interface {
	SendPush(message string) error
}

// ================================================================
// Implementations
// ================================================================

// Email service only supports email
type EmailService struct{}

func (e *EmailService) SendEmail(
	message string,
) error {

	fmt.Printf(
		"[EMAIL] %s\n",
		message,
	)

	return nil
}

// SMS service only supports SMS
type SMSService struct{}

func (s *SMSService) SendSMS(
	message string,
) error {

	fmt.Printf(
		"[SMS] %s\n",
		message,
	)

	return nil
}

// Push service only supports push notifications
type PushService struct{}

func (p *PushService) SendPush(
	message string,
) error {

	fmt.Printf(
		"[PUSH] %s\n",
		message,
	)

	return nil
}

// ================================================================
// Clients
// ================================================================

// Marketing service only depends on email behavior.
type MarketingService struct {
	emailSender EmailSender
}

func NewMarketingService(
	emailSender EmailSender,
) *MarketingService {

	return &MarketingService{
		emailSender: emailSender,
	}
}

func (m *MarketingService) SendCampaign() error {
	return m.emailSender.SendEmail(
		"Big Billion Day Sale!",
	)
}

// OTP service only depends on SMS behavior.
type OTPService struct {
	smsSender SMSSender
}

func NewOTPService(
	smsSender SMSSender,
) *OTPService {

	return &OTPService{
		smsSender: smsSender,
	}
}

func (o *OTPService) SendOTP() error {
	return o.smsSender.SendSMS(
		"Your OTP is 123456",
	)
}

// ================================================================
// Main
// ================================================================

func main() {

	emailService := &EmailService{}
	smsService := &SMSService{}

	marketing :=
		NewMarketingService(emailService)

	otp :=
		NewOTPService(smsService)

	marketing.SendCampaign()
	otp.SendOTP()
}
