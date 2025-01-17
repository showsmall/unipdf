//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package common contains common properties used by the subpackages.
package common ;import (_a "fmt";_c "io";_b "os";_ed "path/filepath";_d "runtime";_ae "time";);

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};const Version ="\u0033\u002e\u0032\u0038\u002e\u0030";func (_dbe WriterLogger )logToWriter (_cb _c .Writer ,_ccf string ,_ga string ,_dfg ...interface{}){_eae (_cb ,_ccf ,_ga ,_dfg );};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _c .Writer )*WriterLogger {_bec :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_bec ;};

// Debug logs debug message.
func (_df ConsoleLogger )Debug (format string ,args ...interface{}){if _df .LogLevel >=LogLevelDebug {_bca :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_df .output (_b .Stdout ,_bca ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;
LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };const _abf =30;

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Notice logs notice message.
func (_bce WriterLogger )Notice (format string ,args ...interface{}){if _bce .LogLevel >=LogLevelNotice {_gg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_bce .logToWriter (_bce .Output ,_gg ,format ,args ...);};};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Info logs info message.
func (_aa ConsoleLogger )Info (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelInfo {_cc :="\u005bI\u004e\u0046\u004f\u005d\u0020";_aa .output (_b .Stdout ,_cc ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Warning logs warning message.
func (_aea ConsoleLogger )Warning (format string ,args ...interface{}){if _aea .LogLevel >=LogLevelWarning {_ea :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_aea .output (_b .Stdout ,_ea ,format ,args ...);};};

// Trace logs trace message.
func (_aeg WriterLogger )Trace (format string ,args ...interface{}){if _aeg .LogLevel >=LogLevelTrace {_eeb :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_aeg .logToWriter (_aeg .Output ,_eeb ,format ,args ...);};};const _aaf ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";
const _ecf =2021;

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};const _fcd =8;func _eae (_gf _c .Writer ,_dgf string ,_da string ,_dfb ...interface{}){_ ,_dc ,_bg ,_ecg :=_d .Caller (3);if !_ecg {_dc ="\u003f\u003f\u003f";_bg =0;}else {_dc =_ed .Base (_dc );
};_ac :=_a .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_dgf ,_dc ,_bg )+_da +"\u000a";_a .Fprintf (_gf ,_ac ,_dfb ...);};var Log Logger =DummyLogger {};

// Error logs error message.
func (_efg WriterLogger )Error (format string ,args ...interface{}){if _efg .LogLevel >=LogLevelError {_fc :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_efg .logToWriter (_efg .Output ,_fc ,format ,args ...);};};

// Debug logs debug message.
func (_abb WriterLogger )Debug (format string ,args ...interface{}){if _abb .LogLevel >=LogLevelDebug {_fd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_abb .logToWriter (_abb .Output ,_fd ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fe WriterLogger )IsLogLevel (level LogLevel )bool {return _fe .LogLevel >=level };

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _ae .Time )string {return t .Format (_aaf )+"\u0020\u0055\u0054\u0043"};

// Trace logs trace message.
func (_cg ConsoleLogger )Trace (format string ,args ...interface{}){if _cg .LogLevel >=LogLevelTrace {_be :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_cg .output (_b .Stdout ,_be ,format ,args ...);};};func (_fa ConsoleLogger )output (_bf _c .Writer ,_aff string ,_bcg string ,_dbgd ...interface{}){_eae (_bf ,_aff ,_bcg ,_dbgd ...);
};

// Notice logs notice message.
func (_dbg ConsoleLogger )Notice (format string ,args ...interface{}){if _dbg .LogLevel >=LogLevelNotice {_de :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_dbg .output (_b .Stdout ,_de ,format ,args ...);};};const _adb =15;

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Error logs error message.
func (_ab ConsoleLogger )Error (format string ,args ...interface{}){if _ab .LogLevel >=LogLevelError {_dd :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ab .output (_b .Stdout ,_dd ,format ,args ...);};};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_edg string ,_db ...interface{});Warning (_f string ,_fg ...interface{});Notice (_ec string ,_g ...interface{});Info (_gb string ,_dg ...interface{});Debug (_dba string ,_eb ...interface{});Trace (_bc string ,_ef ...interface{});
IsLogLevel (_cf LogLevel )bool ;};

// DummyLogger does nothing.
type DummyLogger struct{};

// Warning logs warning message.
func (_ad WriterLogger )Warning (format string ,args ...interface{}){if _ad .LogLevel >=LogLevelWarning {_fee :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ad .logToWriter (_ad .Output ,_fee ,format ,args ...);};};

// Info logs info message.
func (_bfg WriterLogger )Info (format string ,args ...interface{}){if _bfg .LogLevel >=LogLevelInfo {_ee :="\u005bI\u004e\u0046\u004f\u005d\u0020";_bfg .logToWriter (_bfg .Output ,_ee ,format ,args ...);};};const _cce =12;

// LogLevel is the verbosity level for logging.
type LogLevel int ;var ReleasedAt =_ae .Date (_ecf ,_fcd ,_cce ,_adb ,_abf ,0,0,_ae .UTC );

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_af ConsoleLogger )IsLogLevel (level LogLevel )bool {return _af .LogLevel >=level };

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _c .Writer ;};