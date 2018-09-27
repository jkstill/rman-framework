package general

// Standard imports

import "flag"

// Local imports

import "github.com/daviesluke/logger"
import "github.com/daviesluke/setup"
import "github.com/daviesluke/utils"

// local Variables

// Initialise long flags

var configFile = flag.String("config"     , "", "Config File Name")
var database   = flag.String("db"         , "", "Database name")
var errorEmail = flag.String("erroremail" , "", "E-mail list for errors")
var email      = flag.String("email"      , "", "E-mail list for success / error")
var lock       = flag.String("lock"       , "", "Lock name")
var logDir     = flag.String("log"        , "", "Directory for logs")
var resource   = flag.String("resource"   , "", "Resource name")


// Local functions

func init() {
	//
	// Setting up short flags
	//

	flag.StringVar(configFile, "c", "", "Config File Name")
	flag.StringVar(database  , "d", "", "Database name")
	flag.StringVar(errorEmail, "e", "", "E-mail list for errors")
	flag.StringVar(email     , "E", "", "E-mail List for success / error")
	flag.StringVar(lock      , "l", "", "Lock name")
	flag.StringVar(logDir    , "L", "", "Alternative Log directory")
	flag.StringVar(resource  , "r", "", "Resource name")
}

// Global functions

func ValidateFlags () {
	logger.Info("Validating command line arguments ...")

	flag.Parse()

	visitor := func(flagParam *flag.Flag) {
		logger.Infof("Parameter %s set to %s", flagParam.Usage, flagParam.Value)

		// Validate and do appropriate changes for each option

		emailRegEx := "^(([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5}){1,25})+([;.](([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5}){1,25})+)*$"

		if flagParam.Name == "config" || flagParam.Name == "c" {
			logger.Infof("Config file switched from default to named file -> %s", *configFile)
			setup.SwitchConfigFile(*configFile)
			logger.Tracef("Config file now set to %s", setup.ConfigFileName)
		} else if flagParam.Name == "log" || flagParam.Name == "l" {
			logger.Infof("Switching log directory to new directory -> %s", *logDir)
			setup.SwitchLogDir(*logDir)
			//logger.SwitchLog()
			logger.Tracef("Log file now set to %s", setup.LogFileName)
		} else if flagParam.Name == "erroremail" || flagParam.Name == "e" {
			logger.Info("Validating error e-mail addresses ...")
			if utils.CheckRegEx(*errorEmail, emailRegEx) {
				logger.Debugf("Error E-mail - %s validated", *errorEmail)
				setup.SetErrorEmail(*errorEmail)
			} else {
				logger.Errorf("Invalid error e-mail address list - %s", *errorEmail)
			}
		} else if flagParam.Name == "email" || flagParam.Name == "E" {
			logger.Info("Validating e-mail addresses ...")
			if utils.CheckRegEx(*email, emailRegEx) {
				logger.Debugf("E-mail - %s validated", *email)
				setup.SetEmail(*email)
			} else {
				logger.Errorf("Invalid e-mail address list - %s", *email)
			}
		} else if flagParam.Name == "resource" || flagParam.Name == "r" {
			logger.Info("Validating resources ...")
			resourceRegEx := "^([a-zA-Z0-9_]+=[0-9]+)+([:;,.][a-zA-Z0-9_]+=[0-9]+)*$"
			if utils.CheckRegEx(*resource, resourceRegEx) {
				logger.Debugf("Resources - %s - validated", *resource)
				setup.SetResource(*resource)
			} else {
				logger.Errorf("Invalid resources - %s", *resource)
			}
		} else if flagParam.Name == "db" || flagParam.Name == "d" {
			setup.SetDatabase(*database)
		}
	}

	flag.Visit(visitor)

	logger.Info("Process complete")
}

func Cleanup() {
	logger.Infof("Running cleanup ...")

	logger.Infof("Process complete")
}

