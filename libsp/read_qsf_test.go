package libsp

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadQsfMetadata(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}
	if s.Title != "Test survey" {
		t.Errorf("Title = '%s'; want 'Test survey'", s.Title)
	}
	if s.Status != "Inactive" {
		t.Errorf("Status = '%s'; want 'Inactive'", s.Status)
	}
	if s.Description != "Test description" {
		t.Errorf("Description = '%s'; want 'Test description'", s.Description)
	}
	if s.CreatedOn.String() != "2019-02-10 21:50:52 +0000 UTC" {
		t.Errorf("CreatedOn = '%s'; want '2019-02-10 21:50:52 +0000 UTC'", s.CreatedOn)
	}
	if s.ModifiedOn.String() != "2019-02-10 21:55:31 +0000 UTC" {
		t.Errorf("ModifiedOn = '%s'; want '2019-02-10 21:55:31 +0000 UTC'", s.ModifiedOn)
	}
	if s.LaunchedOn.String() != "2019-08-01 01:23:45 +0000 UTC" {
		t.Errorf("LaunchedOn = '%s'; want '2019-08-01 01:23:45 +0000 UTC'", s.LaunchedOn)
	}
}

func TestReadQsfQuestions(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}
	if len(s.Questions) != 10 {
		t.Errorf("len(Questions) = %d; want 10", len(s.Questions))
	}
	if _, ok := s.Questions["QID1"]; !ok {
		t.Error("QID1 not found in s.Questions")
	}
	if _, ok := s.Questions["QID2"]; !ok {
		t.Error("QID2 not found in s.Questions")
	}
	if _, ok := s.Questions["QID3"]; !ok {
		t.Error("QID3 not found in s.Questions")
	}
	if _, ok := s.Questions["QID4"]; !ok {
		t.Error("QID4 not found in s.Questions")
	}
	if _, ok := s.Questions["QID5"]; !ok {
		t.Error("QID5 not found in s.Questions")
	}
	if _, ok := s.Questions["QID6"]; !ok {
		t.Error("QID6 not found in s.Questions")
	}
	if _, ok := s.Questions["QID7"]; !ok {
		t.Error("QID7 not found in s.Questions")
	}
	if _, ok := s.Questions["QID8"]; !ok {
		t.Error("QID8 not found in s.Questions")
	}
	if _, ok := s.Questions["QID9"]; !ok {
		t.Error("QID9 not found in s.Questions")
	}
	if _, ok := s.Questions["QID10"]; !ok {
		t.Error("QID10 not found in s.Questions")
	}
	if _, ok := s.Questions["QID11"]; ok {
		t.Error("QID11 found in s.Questions")
	}
}

func TestIncompleteQSF(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfContentIncomplete))
	s, err := ReadQsf(r)
	if err == nil {
		t.Error("err = nil; want err != nil")
	}
	if s != nil {
		t.Error("survey != nil; want survey = nil")
	}
	if err.Error() != "could not parse: json had no SurveyEntry object" {
		t.Errorf("err = '%s'; want 'could not parse: json had no SurveyEntry object'", err)
	}
}

func TestNilReader(t *testing.T) {
	s, err := ReadQsf(nil)
	if err == nil {
		t.Error("err = nil; want err != nil")
	}
	if err.Error() != "r cannot be nil" {
		t.Errorf("err.Error() = '%s'; want 'r cannot be nil'", err)
	}
	if s != nil {
		t.Error("got != nil; want got = nil")
	}
}

var qsfContent = `{
    "SurveyEntry": {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "SurveyName": "Test survey",
        "SurveyDescription": "Test description",
        "SurveyOwnerID": "UR_5AXpopyMdC5gltr",
        "SurveyBrandID": "google",
        "DivisionID": null,
        "SurveyLanguage": "EN",
        "SurveyActiveResponseSet": "RS_2uzQg6GtgfLcVb7",
        "SurveyStatus": "Inactive",
        "SurveyStartDate": "2019-08-01 01:23:45",
        "SurveyExpirationDate": "0000-00-00 00:00:00",
        "SurveyCreationDate": "2019-02-10 21:50:52",
        "CreatorID": "UR_5AXpopyMdC5gltr",
        "LastModified": "2019-02-10 21:55:31",
        "LastAccessed": "0000-00-00 00:00:00",
        "LastActivated": "0000-00-00 00:00:00",
        "Deleted": null
    },
    "SurveyElements": [
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "BL",
            "PrimaryAttribute": "Survey Blocks",
            "SecondaryAttribute": null,
            "TertiaryAttribute": null,
            "Payload": [
                {
                    "Type": "Default",
                    "Description": "Multiple choice",
                    "ID": "BL_86vwFSQoawhxvMx",
                    "BlockElements": [
                        {
                            "Type": "Question",
                            "QuestionID": "QID1"
                        },
                        {
                            "Type": "Question",
                            "QuestionID": "QID3"
                        },
                        {
                            "Type": "Question",
                            "QuestionID": "QID4"
                        }
                    ]
                },
                {
                    "Type": "Trash",
                    "Description": "Trash / Unused Questions",
                    "ID": "BL_3lqif1AXkIaT2jX",
                    "BlockElements": [
                        {
                            "Type": "Question",
                            "QuestionID": "QID2"
                        }
                    ]
                },
                {
                    "Type": "Standard",
                    "SubType": "",
                    "Description": "Block 1",
                    "ID": "BL_4JaWHcVfXC1Sted",
                    "BlockElements": [
                        {
                            "Type": "Question",
                            "QuestionID": "QID5"
                        },
                        {
                            "Type": "Question",
                            "QuestionID": "QID6"
                        }
                    ]
                },
                {
                    "Type": "Standard",
                    "SubType": "",
                    "Description": "Block 2",
                    "ID": "BL_cZmBNWuCEsNWDPf",
                    "BlockElements": [
                        {
                            "Type": "Question",
                            "QuestionID": "QID7"
                        },
                        {
                            "Type": "Question",
                            "QuestionID": "QID8"
                        },
                        {
                            "Type": "Question",
                            "QuestionID": "QID9"
                        }
                    ]
                },
                {
                    "Type": "Standard",
                    "SubType": "",
                    "Description": "Block 3",
                    "ID": "BL_5hYxuIwD9MN62JT",
                    "BlockElements": [
                        {
                            "Type": "Question",
                            "QuestionID": "QID10"
                        }
                    ]
                }
            ]
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "FL",
            "PrimaryAttribute": "Survey Flow",
            "SecondaryAttribute": null,
            "TertiaryAttribute": null,
            "Payload": {
                "Flow": [
                    {
                        "ID": "BL_86vwFSQoawhxvMx",
                        "Type": "Block",
                        "FlowID": "FL_2"
                    },
                    {
                        "ID": "BL_4JaWHcVfXC1Sted",
                        "Type": "Standard",
                        "FlowID": "FL_3"
                    },
                    {
                        "ID": "BL_cZmBNWuCEsNWDPf",
                        "Type": "Standard",
                        "FlowID": "FL_4"
                    },
                    {
                        "ID": "BL_5hYxuIwD9MN62JT",
                        "Type": "Standard",
                        "FlowID": "FL_5"
                    }
                ],
                "Properties": {
                    "Count": 5
                },
                "FlowID": "FL_1",
                "Type": "Root"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SO",
            "PrimaryAttribute": "Survey Options",
            "SecondaryAttribute": null,
            "TertiaryAttribute": null,
            "Payload": {
                "BackButton": "false",
                "SaveAndContinue": "true",
                "SurveyProtection": "PublicSurvey",
                "BallotBoxStuffingPrevention": "false",
                "NoIndex": "Yes",
                "SecureResponseFiles": "true",
                "SurveyExpiration": "None",
                "SurveyTermination": "DefaultMessage",
                "Header": "",
                "Footer": "",
                "ProgressBarDisplay": "None",
                "PartialData": "+1 week",
                "ValidationMessage": "",
                "PreviousButton": " ← ",
                "NextButton": " → ",
                "SkinLibrary": "google",
                "SkinType": "MQ",
                "Skin": "google_11",
                "NewScoring": 1
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SCO",
            "PrimaryAttribute": "Scoring",
            "SecondaryAttribute": null,
            "TertiaryAttribute": null,
            "Payload": {
                "ScoringCategories": [],
                "ScoringCategoryGroups": [],
                "ScoringSummaryCategory": null,
                "ScoringSummaryAfterQuestions": 0,
                "ScoringSummaryAfterSurvey": 0,
                "DefaultScoringCategory": null,
                "AutoScoringCategory": null
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "PROJ",
            "PrimaryAttribute": "CORE",
            "SecondaryAttribute": null,
            "TertiaryAttribute": "1.1.0",
            "Payload": {
                "ProjectCategory": "CORE",
                "SchemaVersion": "1.1.0"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "STAT",
            "PrimaryAttribute": "Survey Statistics",
            "SecondaryAttribute": null,
            "TertiaryAttribute": null,
            "Payload": {
                "MobileCompatible": false,
                "ID": "Survey Statistics"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "QC",
            "PrimaryAttribute": "Survey Question Count",
            "SecondaryAttribute": "10",
            "TertiaryAttribute": null,
            "Payload": null
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID2",
            "SecondaryAttribute": "Click to write the question text",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Click to write the question text",
                "DefaultChoices": false,
                "DataExportTag": "Q2",
                "QuestionType": "MC",
                "Selector": "MSB",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText"
                },
                "QuestionDescription": "Click to write the question text",
                "Choices": {
                    "1": {
                        "Display": "Click to write Choice 1"
                    },
                    "2": {
                        "Display": "Click to write Choice 2"
                    },
                    "3": {
                        "Display": "Click to write Choice 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID2"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "RS",
            "PrimaryAttribute": "RS_2uzQg6GtgfLcVb7",
            "SecondaryAttribute": "Default Response Set",
            "TertiaryAttribute": null,
            "Payload": null
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID9",
            "SecondaryAttribute": "Form",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Form",
                "DefaultChoices": false,
                "DataExportTag": "Q9",
                "QuestionType": "TE",
                "Selector": "FORM",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText",
                    "AllowFreeResponse": "false"
                },
                "QuestionDescription": "Form",
                "Choices": {
                    "1": {
                        "Display": "Click to write Choice 1"
                    },
                    "2": {
                        "Display": "Click to write Choice 2"
                    },
                    "3": {
                        "Display": "Click to write Choice 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": null
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID9"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID5",
            "SecondaryAttribute": "Matrix",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Matrix",
                "DefaultChoices": false,
                "DataExportTag": "Q5",
                "QuestionType": "Matrix",
                "Selector": "Likert",
                "SubSelector": "SingleAnswer",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText",
                    "TextPosition": "inline",
                    "ChoiceColumnWidth": 25,
                    "RepeatHeaders": "none",
                    "WhiteSpace": "OFF",
                    "MobileFirst": true
                },
                "QuestionDescription": "Matrix",
                "Choices": {
                    "1": {
                        "Display": "Click to write Statement 1"
                    },
                    "2": {
                        "Display": "Click to write Statement 2"
                    },
                    "3": {
                        "Display": "Click to write Statement 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 4,
                "Answers": {
                    "1": {
                        "Display": "Click to write Scale point 1"
                    },
                    "2": {
                        "Display": "Click to write Scale point 2"
                    },
                    "3": {
                        "Display": "Click to write Scale point 3"
                    }
                },
                "AnswerOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "ChoiceDataExportTags": false,
                "QuestionID": "QID5"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID6",
            "SecondaryAttribute": "MaxDiff",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "MaxDiff",
                "DefaultChoices": false,
                "DataExportTag": "Q6",
                "QuestionType": "Matrix",
                "Selector": "MaxDiff",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText",
                    "ChoiceColumnWidth": 25,
                    "MobileFirst": true
                },
                "QuestionDescription": "MaxDiff",
                "Choices": {
                    "1": {
                        "Display": "Click to write Statement 1",
                        "ExclusiveAnswer": true
                    },
                    "2": {
                        "Display": "Click to write Statement 2",
                        "ExclusiveAnswer": true
                    },
                    "3": {
                        "Display": "Click to write Statement 3",
                        "ExclusiveAnswer": true
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 3,
                "Answers": {
                    "1": {
                        "Display": "Click to write Scale point 1",
                        "ExclusiveAnswer": true
                    },
                    "2": {
                        "Display": "Click to write Scale point 2",
                        "ExclusiveAnswer": true
                    }
                },
                "AnswerOrder": [
                    1,
                    2
                ],
                "ChoiceDataExportTags": false,
                "QuestionID": "QID6"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID8",
            "SecondaryAttribute": "Multi line text entry",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Multi line text entry",
                "DefaultChoices": false,
                "DataExportTag": "Q8",
                "QuestionType": "TE",
                "Selector": "ML",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText",
                    "InputWidth": 575,
                    "InputHeight": 167,
                    "AllowFreeResponse": "false"
                },
                "QuestionDescription": "Multi line text entry",
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID8"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID4",
            "SecondaryAttribute": "Multiple answer",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Multiple answer",
                "DataExportTag": "Q4",
                "QuestionType": "MC",
                "Selector": "MAVR",
                "SubSelector": "TX",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText"
                },
                "QuestionDescription": "Multiple answer",
                "Choices": {
                    "1": {
                        "Display": "Click to write Choice 1"
                    },
                    "2": {
                        "Display": "Click to write Choice 2"
                    },
                    "3": {
                        "Display": "Click to write Choice 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID4"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID10",
            "SecondaryAttribute": "Rank order",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Rank order",
                "DefaultChoices": false,
                "DataExportTag": "Q10",
                "QuestionType": "RO",
                "Selector": "DND",
                "SubSelector": "TX",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText"
                },
                "QuestionDescription": "Rank order",
                "Choices": {
                    "1": {
                        "Display": "Click to write Item 1"
                    },
                    "2": {
                        "Display": "Click to write Item 2"
                    },
                    "3": {
                        "Display": "Click to write Item 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID10"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID3",
            "SecondaryAttribute": "Select box",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Select box",
                "DefaultChoices": false,
                "DataExportTag": "Q3",
                "QuestionType": "MC",
                "Selector": "SB",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText"
                },
                "QuestionDescription": "Select box",
                "Choices": {
                    "1": {
                        "Display": "Click to write Choice 1"
                    },
                    "2": {
                        "Display": "Click to write Choice 2"
                    },
                    "3": {
                        "Display": "Click to write Choice 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID3"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID1",
            "SecondaryAttribute": "Single answer",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Single answer",
                "DataExportTag": "Q1",
                "QuestionType": "MC",
                "Selector": "SAVR",
                "SubSelector": "TX",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText"
                },
                "QuestionDescription": "Single answer",
                "Choices": {
                    "1": {
                        "Display": "Click to write Choice 1"
                    },
                    "2": {
                        "Display": "Click to write Choice 2"
                    },
                    "3": {
                        "Display": "Click to write Choice 3"
                    }
                },
                "ChoiceOrder": [
                    "1",
                    "2",
                    "3"
                ],
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID1"
            }
        },
        {
            "SurveyID": "SV_6mudEEycYo5zehT",
            "Element": "SQ",
            "PrimaryAttribute": "QID7",
            "SecondaryAttribute": "Single line text entry",
            "TertiaryAttribute": null,
            "Payload": {
                "QuestionText": "Single line text entry",
                "DefaultChoices": false,
                "DataExportTag": "Q7",
                "QuestionType": "TE",
                "Selector": "SL",
                "Configuration": {
                    "QuestionDescriptionOption": "UseText",
                    "AllowFreeResponse": "false"
                },
                "QuestionDescription": "Single line text entry",
                "Validation": {
                    "Settings": {
                        "ForceResponse": "OFF",
                        "ForceResponseType": "ON",
                        "Type": "None"
                    }
                },
                "GradingData": [],
                "Language": [],
                "NextChoiceId": 4,
                "NextAnswerId": 1,
                "QuestionID": "QID7"
            }
        }
    ]
}`

var qsfContentIncomplete = `{}`
