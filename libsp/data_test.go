package libsp

// spell-checker: disable

var qsfTestContent = `{
    "SurveyEntry": {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "SurveyName": "Test survey",
        "SurveyDescription": "Test description",
        "SurveyOwnerID": "UR_5AXpopyMdC5gltr",
        "SurveyBrandID": "google",
        "DivisionID": null,
        "SurveyLanguage": "EN",
        "SurveyActiveResponseSet": "RS_2uzQg6GtgfLcVb7",
        "SurveyStatus": "Active",
        "SurveyStartDate": "2019-08-01 01:23:45",
        "SurveyExpirationDate": "0000-00-00 00:00:00",
        "SurveyCreationDate": "2019-02-10 21:50:52",
        "CreatorID": "UR_5AXpopyMdC5gltr",
        "LastModified": "2019-08-20 12:23:35",
        "LastAccessed": "0000-00-00 00:00:00",
        "LastActivated": "2019-08-20 12:19:02",
        "Deleted": null
    },
    "SurveyElements": [{
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "BL",
        "PrimaryAttribute": "Survey Blocks",
        "SecondaryAttribute": null,
        "TertiaryAttribute": null,
        "Payload": [{
            "Type": "Default",
            "Description": "Multiple choice",
            "ID": "BL_86vwFSQoawhxvMx",
            "BlockElements": [{
                "Type": "Question",
                "QuestionID": "QID14"
            }, {
                "Type": "Question",
                "QuestionID": "QID1"
            }, {
                "Type": "Question",
                "QuestionID": "QID3"
            }, {
                "Type": "Question",
                "QuestionID": "QID4"
            }, {
                "Type": "Question",
                "QuestionID": "QID11"
            }]
        }, {
            "Type": "Trash",
            "Description": "Trash \/ Unused Questions",
            "ID": "BL_3lqif1AXkIaT2jX",
            "BlockElements": [{
                "Type": "Question",
                "QuestionID": "QID2"
            }, {
                "Type": "Question",
                "QuestionID": "QID12"
            }]
        }, {
            "Type": "Standard",
            "SubType": "",
            "Description": "Block 1",
            "ID": "BL_4JaWHcVfXC1Sted",
            "BlockElements": [{
                "Type": "Question",
                "QuestionID": "QID5"
            }, {
                "Type": "Question",
                "QuestionID": "QID13"
            }, {
                "Type": "Question",
                "QuestionID": "QID6"
            }, {
                "Type": "Question",
                "QuestionID": "QID16"
            }]
        }, {
            "Type": "Standard",
            "SubType": "",
            "Description": "Block 2",
            "ID": "BL_cZmBNWuCEsNWDPf",
            "BlockElements": [{
                "Type": "Question",
                "QuestionID": "QID7"
            }, {
                "Type": "Question",
                "QuestionID": "QID8"
            }, {
                "Type": "Question",
                "QuestionID": "QID9"
            }]
        }, {
            "Type": "Standard",
            "SubType": "",
            "Description": "Block 3",
            "ID": "BL_5hYxuIwD9MN62JT",
            "BlockElements": [{
                "Type": "Question",
                "QuestionID": "QID10"
            }, {
                "Type": "Question",
                "QuestionID": "QID17"
            }, {
                "Type": "Question",
                "QuestionID": "QID15"
            }]
        }]
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "FL",
        "PrimaryAttribute": "Survey Flow",
        "SecondaryAttribute": null,
        "TertiaryAttribute": null,
        "Payload": {
            "Flow": [{
                "ID": "BL_86vwFSQoawhxvMx",
                "Type": "Block",
                "FlowID": "FL_2"
            }, {
                "ID": "BL_4JaWHcVfXC1Sted",
                "Type": "Standard",
                "FlowID": "FL_3"
            }, {
                "ID": "BL_cZmBNWuCEsNWDPf",
                "Type": "Standard",
                "FlowID": "FL_4"
            }, {
                "ID": "BL_5hYxuIwD9MN62JT",
                "Type": "Standard",
                "FlowID": "FL_5"
            }],
            "Properties": {
                "Count": 5
            },
            "FlowID": "FL_1",
            "Type": "Root"
        }
    }, {
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
            "PreviousButton": " \u2190 ",
            "NextButton": " \u2192 ",
            "SkinLibrary": "google",
            "SkinType": "MQ",
            "Skin": "google_11",
            "NewScoring": 1
        }
    }, {
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
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "PROJ",
        "PrimaryAttribute": "CORE",
        "SecondaryAttribute": null,
        "TertiaryAttribute": "1.1.0",
        "Payload": {
            "ProjectCategory": "CORE",
            "SchemaVersion": "1.1.0"
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "STAT",
        "PrimaryAttribute": "Survey Statistics",
        "SecondaryAttribute": null,
        "TertiaryAttribute": null,
        "Payload": {
            "MobileCompatible": false,
            "ID": "Survey Statistics"
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "QC",
        "PrimaryAttribute": "Survey Question Count",
        "SecondaryAttribute": "17",
        "TertiaryAttribute": null,
        "Payload": null
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID99",
        "SecondaryAttribute": "Browser Meta Info",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Browser Meta Info",
            "DefaultChoices": false,
            "QuestionType": "Meta",
            "Selector": "Browser",
            "Configuration": {
                "QuestionDescriptionOption": "UseText"
            },
            "QuestionDescription": "Browser Meta Info",
            "Choices": {
                "1": {
                    "Display": "Browser",
                    "TextEntry": 1
                },
                "2": {
                    "Display": "Version",
                    "TextEntry": 1
                },
                "3": {
                    "Display": "Operating System",
                    "TextEntry": 1
                },
                "4": {
                    "Display": "Screen Resolution",
                    "TextEntry": 1
                },
                "5": {
                    "Display": "Flash Version",
                    "TextEntry": 1
                },
                "6": {
                    "Display": "Java Support",
                    "TextEntry": 1
                },
                "7": {
                    "Display": "User Agent",
                    "TextEntry": 1
                }
            },
            "GradingData": [],
            "Language": [],
            "QuestionText_Unsafe": "Browser Meta Info",
            "DataExportTag": "Q1.2",
            "QuestionID": "QID99"
        }
    }, {
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
            "ChoiceOrder": ["1", "2", "3"],
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
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID12",
        "SecondaryAttribute": "Click to write the question text",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Click to write the question text",
            "DataExportTag": "Q12",
            "QuestionType": "MC",
            "Selector": "SAVR",
            "SubSelector": "TX",
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
            "ChoiceOrder": ["1", "2", "3"],
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
            "QuestionID": "QID12"
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "RS",
        "PrimaryAttribute": "RS_2uzQg6GtgfLcVb7",
        "SecondaryAttribute": "Default Response Set",
        "TertiaryAttribute": null,
        "Payload": null
    }, {
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
                    "Display": "Form line 1"
                },
                "2": {
                    "Display": "Form line 2"
                },
                "3": {
                    "Display": "Form line 3"
                }
            },
            "ChoiceOrder": ["1", "2", "3"],
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
            "QuestionID": "QID9",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            },
            "SearchSource": {
                "AllowFreeResponse": "false"
            }
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID13",
        "SecondaryAttribute": "Matrix multiple response per row",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Matrix multiple response per row",
            "DefaultChoices": false,
            "DataExportTag": "Q13",
            "QuestionID": "QID13",
            "QuestionType": "Matrix",
            "Selector": "Likert",
            "SubSelector": "MultipleAnswer",
            "Configuration": {
                "QuestionDescriptionOption": "UseText",
                "TextPosition": "inline",
                "ChoiceColumnWidth": 25,
                "RepeatHeaders": "none",
                "WhiteSpace": "OFF",
                "MobileFirst": true
            },
            "QuestionDescription": "Matrix multiple response per row",
            "Choices": {
                "1": {
                    "Display": "Row 1"
                },
                "2": {
                    "Display": "Row 2"
                },
                "3": {
                    "Display": "Row 3"
                },
                "4": {
                    "Display": "Other:",
                    "TextEntry": "true"
                },
                "5": {
                    "Display": ""
                }
            },
            "ChoiceOrder": ["1", "2", "3", "5", "4"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 7,
            "NextAnswerId": 4,
            "Answers": {
                "1": {
                    "Display": "Col 1"
                },
                "2": {
                    "Display": "Col 2"
                },
                "3": {
                    "Display": "Col 3"
                }
            },
            "AnswerOrder": ["1", "2", "3"],
            "ChoiceDataExportTags": false
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID5",
        "SecondaryAttribute": "Matrix single response per row",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Matrix single response per row",
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
            "QuestionDescription": "Matrix single response per row",
            "Choices": {
                "1": {
                    "Display": "Click to write Statement 1"
                },
                "2": {
                    "Display": "Click to write Statement 2"
                },
                "3": {
                    "Display": "Click to write Statement 3"
                },
                "4": {
                    "Display": "Other:",
                    "TextEntry": "true"
                }
            },
            "ChoiceOrder": ["1", "2", "3", "4"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 5,
            "NextAnswerId": 5,
            "Answers": {
                "1": {
                    "Display": "Click to write Scale point 1"
                },
                "2": {
                    "Display": "Click to write Scale point 2"
                },
                "3": {
                    "Display": "Click to write Scale point 3"
                },
                "4": {
                    "Display": "n\/a"
                }
            },
            "AnswerOrder": ["1", "2", "3", "4"],
            "ChoiceDataExportTags": false,
            "QuestionID": "QID5",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            },
            "AnalyzeChoices": {
                "4": "No"
            }
        }
    }, {
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
            "ChoiceOrder": ["1", "2", "3"],
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
            "AnswerOrder": [1, 2],
            "ChoiceDataExportTags": false,
            "QuestionID": "QID6"
        }
    }, {
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
    }, {
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
                },
                "4": {
                    "Display": "None",
                    "ExclusiveAnswer": true
                },
                "5": {
                    "Display": "Other1",
                    "TextEntry": "true"
                },
                "6": {
                    "Display": "Other2",
                    "TextEntry": "true"
                }
            },
            "ChoiceOrder": ["1", "2", "3", "5", "6", "4"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "Language": [],
            "NextChoiceId": 7,
            "NextAnswerId": 1,
            "QuestionID": "QID4",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            }
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID15",
        "SecondaryAttribute": "On a scale from 0-10, how likely are you to recommend [INSERT COMPANY NAME HERE] to a friend or c...",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "On a scale from 0-10, how likely are you to recommend [INSERT COMPANY NAME HERE] to a friend or colleague?",
            "DefaultChoices": false,
            "DataExportTag": "Q15",
            "QuestionType": "MC",
            "Selector": "NPS",
            "Configuration": {
                "QuestionDescriptionOption": "UseText"
            },
            "QuestionDescription": "On a scale from 0-10, how likely are you to recommend [INSERT COMPANY NAME HERE] to a friend or c...",
            "Choices": [{
                "Display": "0"
            }, {
                "Display": "1"
            }, {
                "Display": "2"
            }, {
                "Display": "3"
            }, {
                "Display": "4"
            }, {
                "Display": "5"
            }, {
                "Display": "6"
            }, {
                "Display": "7"
            }, {
                "Display": "8"
            }, {
                "Display": "9"
            }, {
                "Display": "10"
            }],
            "ChoiceOrder": ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 11,
            "NextAnswerId": 1,
            "ColumnLabels": [{
                "Display": "Not at all likely",
                "IsLabelDefault": true
            }, {
                "Display": "Extremely likely",
                "IsLabelDefault": true
            }],
            "QuestionID": "QID15"
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID11",
        "SecondaryAttribute": "Other text entry",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Other text entry",
            "DefaultChoices": false,
            "DataExportTag": "Q11",
            "QuestionType": "MC",
            "Selector": "SAVR",
            "SubSelector": "TX",
            "Configuration": {
                "QuestionDescriptionOption": "UseText"
            },
            "QuestionDescription": "Other text entry",
            "Choices": {
                "1": {
                    "Display": "Click to write Choice 1 (ordered 2nd)"
                },
                "2": {
                    "Display": "Click to write Choice 2 (ordered 1st)"
                },
                "3": {
                    "Display": "Click to write Choice 3",
                    "TextEntry": "true",
                    "InputHeight": 24,
                    "InputWidth": 284
                }
            },
            "ChoiceOrder": ["2", "1", "3"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 5,
            "NextAnswerId": 4,
            "QuestionID": "QID11",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            }
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID17",
        "SecondaryAttribute": "Pick, group, rank",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Pick, group, rank",
            "DefaultChoices": false,
            "DataExportTag": "Q17",
            "QuestionType": "PGR",
            "Selector": "DragAndDrop",
            "SubSelector": "NoColumns",
            "Configuration": {
                "QuestionDescriptionOption": "UseText",
                "Stack": false,
                "StackItemsInGroups": false
            },
            "QuestionDescription": "Pick, group, rank",
            "Choices": {
                "1": {
                    "Display": "Item 1"
                },
                "2": {
                    "Display": "Item 2"
                },
                "3": {
                    "Display": "Item 3"
                },
                "4": {
                    "Display": "Item 4"
                },
                "5": {
                    "Display": "Other:",
                    "TextEntry": "true"
                }
            },
            "ChoiceOrder": ["1", "2", "3", "4", "5"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "Type": "None",
                    "MinChoices": null,
                    "MaxChoices": null
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 6,
            "NextAnswerId": 4,
            "Groups": ["Click to write Group 1", "Click to write Group 2", "Click to write Group 3"],
            "NumberOfGroups": 3,
            "QuestionID": "QID17"
        }
    }, {
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
            "ChoiceOrder": ["1", "2", "3"],
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
            "QuestionID": "QID10",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            }
        }
    }, {
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
            "ChoiceOrder": ["1", "2", "3"],
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
    }, {
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
            "ChoiceOrder": ["1", "2", "3"],
            "Validation": {
                "Settings": {
                    "ForceResponse": "OFF",
                    "ForceResponseType": "ON",
                    "Type": "None"
                }
            },
            "Language": [],
            "NextChoiceId": 5,
            "NextAnswerId": 1,
            "QuestionID": "QID1",
            "DataVisibility": {
                "Private": false,
                "Hidden": false
            }
        }
    }, {
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
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID14",
        "SecondaryAttribute": "Some descriptive text to start things off.",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Some descriptive text to start things off.",
            "DefaultChoices": false,
            "DataExportTag": "Q14",
            "QuestionType": "DB",
            "Selector": "TB",
            "Configuration": {
                "QuestionDescriptionOption": "UseText"
            },
            "QuestionDescription": "Some descriptive text to start things off.",
            "ChoiceOrder": [],
            "Validation": {
                "Settings": {
                    "Type": "None"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 4,
            "NextAnswerId": 1,
            "QuestionID": "QID14"
        }
    }, {
        "SurveyID": "SV_6mudEEycYo5zehT",
        "Element": "SQ",
        "PrimaryAttribute": "QID16",
        "SecondaryAttribute": "Timing",
        "TertiaryAttribute": null,
        "Payload": {
            "QuestionText": "Timing",
            "DefaultChoices": false,
            "DataExportTag": "Q16",
            "QuestionType": "Timing",
            "Selector": "PageTimer",
            "Configuration": {
                "QuestionDescriptionOption": "UseText",
                "MinSeconds": "0",
                "MaxSeconds": "0"
            },
            "QuestionDescription": "Timing",
            "Choices": {
                "1": {
                    "Display": "First Click"
                },
                "2": {
                    "Display": "Last Click"
                },
                "3": {
                    "Display": "Page Submit"
                },
                "4": {
                    "Display": "Click Count"
                }
            },
            "GradingData": [],
            "Language": [],
            "NextChoiceId": 16,
            "NextAnswerId": 1,
            "QuestionID": "QID16"
        }
    }]
}`

var qsfContentIncomplete = `{}`

var qsfTestContentMin = `{"SurveyEntry":{"SurveyID":"SV_6mudEEycYo5zehT","SurveyName":"Test survey","SurveyDescription":"Test description","SurveyOwnerID":"UR_5AXpopyMdC5gltr","SurveyBrandID":"google","DivisionID":null,"SurveyLanguage":"EN","SurveyActiveResponseSet":"RS_2uzQg6GtgfLcVb7","SurveyStatus":"Inactive","SurveyStartDate":"2019-08-01 01:23:45","SurveyExpirationDate":"0000-00-00 00:00:00","SurveyCreationDate":"2019-02-10 21:50:52","CreatorID":"UR_5AXpopyMdC5gltr","LastModified":"2019-02-10 21:55:31","LastAccessed":"0000-00-00 00:00:00","LastActivated":"0000-00-00 00:00:00","Deleted":null},"SurveyElements":[{"SurveyID":"SV_6mudEEycYo5zehT","Element":"BL","PrimaryAttribute":"Survey Blocks","SecondaryAttribute":null,"TertiaryAttribute":null,"Payload":[{"Type":"Default","Description":"Multiple choice","ID":"BL_86vwFSQoawhxvMx","BlockElements":[{"Type":"Question","QuestionID":"QID1"},{"Type":"Question","QuestionID":"QID3"},{"Type":"Question","QuestionID":"QID4"}]},{"Type":"Trash","Description":"Trash / Unused Questions","ID":"BL_3lqif1AXkIaT2jX","BlockElements":[{"Type":"Question","QuestionID":"QID2"}]},{"Type":"Standard","SubType":"","Description":"Block 1","ID":"BL_4JaWHcVfXC1Sted","BlockElements":[{"Type":"Question","QuestionID":"QID5"},{"Type":"Question","QuestionID":"QID6"}]},{"Type":"Standard","SubType":"","Description":"Block 2","ID":"BL_cZmBNWuCEsNWDPf","BlockElements":[{"Type":"Question","QuestionID":"QID7"},{"Type":"Question","QuestionID":"QID8"},{"Type":"Question","QuestionID":"QID9"}]},{"Type":"Standard","SubType":"","Description":"Block 3","ID":"BL_5hYxuIwD9MN62JT","BlockElements":[{"Type":"Question","QuestionID":"QID10"}]}]},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"FL","PrimaryAttribute":"Survey Flow","SecondaryAttribute":null,"TertiaryAttribute":null,"Payload":{"Flow":[{"ID":"BL_86vwFSQoawhxvMx","Type":"Block","FlowID":"FL_2"},{"ID":"BL_4JaWHcVfXC1Sted","Type":"Standard","FlowID":"FL_3"},{"ID":"BL_cZmBNWuCEsNWDPf","Type":"Standard","FlowID":"FL_4"},{"ID":"BL_5hYxuIwD9MN62JT","Type":"Standard","FlowID":"FL_5"}],"Properties":{"Count":5},"FlowID":"FL_1","Type":"Root"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SO","PrimaryAttribute":"Survey Options","SecondaryAttribute":null,"TertiaryAttribute":null,"Payload":{"BackButton":"false","SaveAndContinue":"true","SurveyProtection":"PublicSurvey","BallotBoxStuffingPrevention":"false","NoIndex":"Yes","SecureResponseFiles":"true","SurveyExpiration":"None","SurveyTermination":"DefaultMessage","Header":"","Footer":"","ProgressBarDisplay":"None","PartialData":"+1 week","ValidationMessage":"","PreviousButton":" ← ","NextButton":" → ","SkinLibrary":"google","SkinType":"MQ","Skin":"google_11","NewScoring":1}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SCO","PrimaryAttribute":"Scoring","SecondaryAttribute":null,"TertiaryAttribute":null,"Payload":{"ScoringCategories":[],"ScoringCategoryGroups":[],"ScoringSummaryCategory":null,"ScoringSummaryAfterQuestions":0,"ScoringSummaryAfterSurvey":0,"DefaultScoringCategory":null,"AutoScoringCategory":null}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"PROJ","PrimaryAttribute":"CORE","SecondaryAttribute":null,"TertiaryAttribute":"1.1.0","Payload":{"ProjectCategory":"CORE","SchemaVersion":"1.1.0"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"STAT","PrimaryAttribute":"Survey Statistics","SecondaryAttribute":null,"TertiaryAttribute":null,"Payload":{"MobileCompatible":false,"ID":"Survey Statistics"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"QC","PrimaryAttribute":"Survey Question Count","SecondaryAttribute":"10","TertiaryAttribute":null,"Payload":null},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID2","SecondaryAttribute":"Click to write the question text","TertiaryAttribute":null,"Payload":{"QuestionText":"Click to write the question text","DefaultChoices":false,"DataExportTag":"Q2","QuestionType":"MC","Selector":"MSB","Configuration":{"QuestionDescriptionOption":"UseText"},"QuestionDescription":"Click to write the question text","Choices":{"1":{"Display":"Click to write Choice 1"},"2":{"Display":"Click to write Choice 2"},"3":{"Display":"Click to write Choice 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID2"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"RS","PrimaryAttribute":"RS_2uzQg6GtgfLcVb7","SecondaryAttribute":"Default Response Set","TertiaryAttribute":null,"Payload":null},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID9","SecondaryAttribute":"Form","TertiaryAttribute":null,"Payload":{"QuestionText":"Form","DefaultChoices":false,"DataExportTag":"Q9","QuestionType":"TE","Selector":"FORM","Configuration":{"QuestionDescriptionOption":"UseText","AllowFreeResponse":"false"},"QuestionDescription":"Form","Choices":{"1":{"Display":"Click to write Choice 1 (ordered third)"},"2":{"Display":"Click to write Choice 2 (ordered second)"},"3":{"Display":"Click to write Choice 3 (ordered first)"}},"ChoiceOrder":["3","2","1"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":null}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID9"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID5","SecondaryAttribute":"Matrix","TertiaryAttribute":null,"Payload":{"QuestionText":"Matrix","DefaultChoices":false,"DataExportTag":"Q5","QuestionType":"Matrix","Selector":"Likert","SubSelector":"SingleAnswer","Configuration":{"QuestionDescriptionOption":"UseText","TextPosition":"inline","ChoiceColumnWidth":25,"RepeatHeaders":"none","WhiteSpace":"OFF","MobileFirst":true},"QuestionDescription":"Matrix","Choices":{"1":{"Display":"Click to write Statement 1"},"2":{"Display":"Click to write Statement 2"},"3":{"Display":"Click to write Statement 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":4,"Answers":{"1":{"Display":"Click to write Scale point 1"},"2":{"Display":"Click to write Scale point 2"},"3":{"Display":"Click to write Scale point 3"}},"AnswerOrder":["1","2","3"],"ChoiceDataExportTags":false,"QuestionID":"QID5"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID6","SecondaryAttribute":"MaxDiff","TertiaryAttribute":null,"Payload":{"QuestionText":"MaxDiff","DefaultChoices":false,"DataExportTag":"Q6","QuestionType":"Matrix","Selector":"MaxDiff","Configuration":{"QuestionDescriptionOption":"UseText","ChoiceColumnWidth":25,"MobileFirst":true},"QuestionDescription":"MaxDiff","Choices":{"1":{"Display":"Click to write Statement 1","ExclusiveAnswer":true},"2":{"Display":"Click to write Statement 2","ExclusiveAnswer":true},"3":{"Display":"Click to write Statement 3","ExclusiveAnswer":true}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":3,"Answers":{"1":{"Display":"Click to write Scale point 1","ExclusiveAnswer":true},"2":{"Display":"Click to write Scale point 2","ExclusiveAnswer":true}},"AnswerOrder":[1,2],"ChoiceDataExportTags":false,"QuestionID":"QID6"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID8","SecondaryAttribute":"Multi line text entry","TertiaryAttribute":null,"Payload":{"QuestionText":"Multi line text entry","DefaultChoices":false,"DataExportTag":"Q8","QuestionType":"TE","Selector":"ML","Configuration":{"QuestionDescriptionOption":"UseText","InputWidth":575,"InputHeight":167,"AllowFreeResponse":"false"},"QuestionDescription":"Multi line text entry","Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID8"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID4","SecondaryAttribute":"Multiple answer","TertiaryAttribute":null,"Payload":{"QuestionText":"Multiple answer","DataExportTag":"Q4","QuestionType":"MC","Selector":"MAVR","SubSelector":"TX","Configuration":{"QuestionDescriptionOption":"UseText"},"QuestionDescription":"Multiple answer","Choices":{"1":{"Display":"Click to write Choice 1"},"2":{"Display":"Click to write Choice 2"},"3":{"Display":"Click to write Choice 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID4"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID10","SecondaryAttribute":"Rank order","TertiaryAttribute":null,"Payload":{"QuestionText":"Rank order","DefaultChoices":false,"DataExportTag":"Q10","QuestionType":"RO","Selector":"DND","SubSelector":"TX","Configuration":{"QuestionDescriptionOption":"UseText"},"QuestionDescription":"Rank order","Choices":{"1":{"Display":"Click to write Item 1"},"2":{"Display":"Click to write Item 2"},"3":{"Display":"Click to write Item 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID10"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID3","SecondaryAttribute":"Select box","TertiaryAttribute":null,"Payload":{"QuestionText":"Select box","DefaultChoices":false,"DataExportTag":"Q3","QuestionType":"MC","Selector":"SB","Configuration":{"QuestionDescriptionOption":"UseText"},"QuestionDescription":"Select box","Choices":{"1":{"Display":"Click to write Choice 1"},"2":{"Display":"Click to write Choice 2"},"3":{"Display":"Click to write Choice 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID3"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID1","SecondaryAttribute":"Single answer","TertiaryAttribute":null,"Payload":{"QuestionText":"Single answer","DataExportTag":"Q1","QuestionType":"MC","Selector":"SAVR","SubSelector":"TX","Configuration":{"QuestionDescriptionOption":"UseText"},"QuestionDescription":"Single answer","Choices":{"1":{"Display":"Click to write Choice 1"},"2":{"Display":"Click to write Choice 2"},"3":{"Display":"Click to write Choice 3"}},"ChoiceOrder":["1","2","3"],"Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID1"}},{"SurveyID":"SV_6mudEEycYo5zehT","Element":"SQ","PrimaryAttribute":"QID7","SecondaryAttribute":"Single line text entry","TertiaryAttribute":null,"Payload":{"QuestionText":"Single line text entry","DefaultChoices":false,"DataExportTag":"Q7","QuestionType":"TE","Selector":"SL","Configuration":{"QuestionDescriptionOption":"UseText","AllowFreeResponse":"false"},"QuestionDescription":"Single line text entry","Validation":{"Settings":{"ForceResponse":"OFF","ForceResponseType":"ON","Type":"None"}},"GradingData":[],"Language":[],"NextChoiceId":4,"NextAnswerId":1,"QuestionID":"QID7"}}]}`

var xmlTestContent = `<?xml version="1.0" ?>
<Responses>
	<Response>
		<startDate>2019-08-20 12:42:28</startDate>
		<endDate>2019-08-20 12:44:31</endDate>
		<status>IP Address</status>
		<ipAddress>*******</ipAddress>
		<progress>100</progress>
		<duration>122</duration>
		<finished>True</finished>
		<recordedDate>2019-08-20 12:44:31</recordedDate>
		<_recordId>R_1dtWhiBDD96nfyk</_recordId>
		<recipientLastName>*******</recipientLastName>
		<recipientFirstName>*******</recipientFirstName>
		<recipientEmail>*******</recipientEmail>
		<externalDataReference>*******</externalDataReference>
		<locationLatitude>*******</locationLatitude>
		<locationLongitude>*******</locationLongitude>
		<distributionChannel>anonymous</distributionChannel>
		<userLanguage>EN</userLanguage>
		<QID1>Click to write Choice 1</QID1>
		<QID3>Click to write Choice 2</QID3>
		<QID4_1></QID4_1>
		<QID4_2></QID4_2>
		<QID4_3>Click to write Choice 3</QID4_3>
		<QID4_5>Other1</QID4_5>
		<QID4_6>Other2</QID4_6>
		<QID4_4></QID4_4>
		<QID4_5_TEXT>other response 1</QID4_5_TEXT>
		<QID4_6_TEXT>other response 2</QID4_6_TEXT>
		<QID11>Click to write Choice 2</QID11>
		<QID11_3_TEXT></QID11_3_TEXT>
		<QID5_1>Click to write Scale point 1</QID5_1>
		<QID5_2>Click to write Scale point 2</QID5_2>
		<QID5_3>Click to write Scale point 3</QID5_3>
		<QID5_4>Click to write Scale point 3</QID5_4>
		<QID5_4_TEXT>other matrix row</QID5_4_TEXT>
		<QID13_1_1>Col 1</QID13_1_1>
		<QID13_1_2>Col 2</QID13_1_2>
		<QID13_1_3></QID13_1_3>
		<QID13_2_1></QID13_2_1>
		<QID13_2_2>Col 2</QID13_2_2>
		<QID13_2_3>Col 3</QID13_2_3>
		<QID13_3_1></QID13_3_1>
		<QID13_3_2>Col 2</QID13_3_2>
		<QID13_3_3></QID13_3_3>
		<QID13_5_1></QID13_5_1>
		<QID13_5_2>Col 2</QID13_5_2>
		<QID13_5_3></QID13_5_3>
		<QID13_4_1></QID13_4_1>
		<QID13_4_2></QID13_4_2>
		<QID13_4_3>Col 3</QID13_4_3>
		<QID13_4_TEXT>other matrix multiple row</QID13_4_TEXT>
		<QID6_1>Click to write Scale point 1</QID6_1>
		<QID6_2></QID6_2>
		<QID6_3>Click to write Scale point 2</QID6_3>
		<QID16_FIRST_CLICK>1.313</QID16_FIRST_CLICK>
		<QID16_LAST_CLICK>32.89</QID16_LAST_CLICK>
		<QID16_PAGE_SUBMIT>33.84</QID16_PAGE_SUBMIT>
		<QID16_CLICK_COUNT>15</QID16_CLICK_COUNT>
		<QID7_TEXT>one line of text</QID7_TEXT>
		<QID8_TEXT>multiple
lines
of
text?</QID8_TEXT>
		<QID9_1>field 1</QID9_1>
		<QID9_2>field 2</QID9_2>
		<QID9_3>field 3</QID9_3>
		<QID10_1>3</QID10_1>
		<QID10_2>2</QID10_2>
		<QID10_3>1</QID10_3>
		<QID17_0_GROUP_1>Item 1</QID17_0_GROUP_1>
		<QID17_0_GROUP_2></QID17_0_GROUP_2>
		<QID17_0_GROUP_3>Item 3</QID17_0_GROUP_3>
		<QID17_0_GROUP_4>Item 4</QID17_0_GROUP_4>
		<QID17_0_GROUP_5></QID17_0_GROUP_5>
		<QID17_1_GROUP_1></QID17_1_GROUP_1>
		<QID17_1_GROUP_2>Item 2</QID17_1_GROUP_2>
		<QID17_1_GROUP_3></QID17_1_GROUP_3>
		<QID17_1_GROUP_4></QID17_1_GROUP_4>
		<QID17_1_GROUP_5></QID17_1_GROUP_5>
		<QID17_2_GROUP_1></QID17_2_GROUP_1>
		<QID17_2_GROUP_2></QID17_2_GROUP_2>
		<QID17_2_GROUP_3></QID17_2_GROUP_3>
		<QID17_2_GROUP_4></QID17_2_GROUP_4>
		<QID17_2_GROUP_5>Other:</QID17_2_GROUP_5>
		<QID17_G0_1_RANK>3</QID17_G0_1_RANK>
		<QID17_G0_2_RANK></QID17_G0_2_RANK>
		<QID17_G0_3_RANK>2</QID17_G0_3_RANK>
		<QID17_G0_4_RANK>1</QID17_G0_4_RANK>
		<QID17_G0_5_RANK></QID17_G0_5_RANK>
		<QID17_G1_1_RANK></QID17_G1_1_RANK>
		<QID17_G1_2_RANK>1</QID17_G1_2_RANK>
		<QID17_G1_3_RANK></QID17_G1_3_RANK>
		<QID17_G1_4_RANK></QID17_G1_4_RANK>
		<QID17_G1_5_RANK></QID17_G1_5_RANK>
		<QID17_G2_1_RANK></QID17_G2_1_RANK>
		<QID17_G2_2_RANK></QID17_G2_2_RANK>
		<QID17_G2_3_RANK></QID17_G2_3_RANK>
		<QID17_G2_4_RANK></QID17_G2_4_RANK>
		<QID17_G2_5_RANK>1</QID17_G2_5_RANK>
		<QID17_5_TEXT>in group 3</QID17_5_TEXT>
		<QID15_NPS_GROUP>Detractor</QID15_NPS_GROUP>
		<QID15>6</QID15>
		<Q_DataPolicyViolations></Q_DataPolicyViolations>
	</Response>
	<Response>
		<startDate>2019-08-20 12:44:50</startDate>
		<endDate>2019-08-20 12:46:35</endDate>
		<status>IP Address</status>
		<ipAddress>*******</ipAddress>
		<progress>100</progress>
		<duration>104</duration>
		<finished>True</finished>
		<recordedDate>2019-08-20 12:46:35</recordedDate>
		<_recordId>R_z72KJQMnr3lxZGp</_recordId>
		<recipientLastName>*******</recipientLastName>
		<recipientFirstName>*******</recipientFirstName>
		<recipientEmail>*******</recipientEmail>
		<externalDataReference>*******</externalDataReference>
		<locationLatitude>*******</locationLatitude>
		<locationLongitude>*******</locationLongitude>
		<distributionChannel>anonymous</distributionChannel>
		<userLanguage>EN</userLanguage>
		<QID1>Click to write Choice 3</QID1>
		<QID3>Click to write Choice 2</QID3>
		<QID4_1></QID4_1>
		<QID4_2></QID4_2>
		<QID4_3></QID4_3>
		<QID4_5></QID4_5>
		<QID4_6></QID4_6>
		<QID4_4>None</QID4_4>
		<QID4_5_TEXT></QID4_5_TEXT>
		<QID4_6_TEXT></QID4_6_TEXT>
		<QID11>Click to write Choice 3</QID11>
		<QID11_3_TEXT>other text</QID11_3_TEXT>
		<QID5_1>Click to write Scale point 3</QID5_1>
		<QID5_2>Click to write Scale point 2</QID5_2>
		<QID5_3></QID5_3>
		<QID5_4></QID5_4>
		<QID5_4_TEXT></QID5_4_TEXT>
		<QID13_1_1>Col 1</QID13_1_1>
		<QID13_1_2></QID13_1_2>
		<QID13_1_3></QID13_1_3>
		<QID13_2_1></QID13_2_1>
		<QID13_2_2></QID13_2_2>
		<QID13_2_3></QID13_2_3>
		<QID13_3_1>Col 1</QID13_3_1>
		<QID13_3_2>Col 2</QID13_3_2>
		<QID13_3_3>Col 3</QID13_3_3>
		<QID13_5_1></QID13_5_1>
		<QID13_5_2></QID13_5_2>
		<QID13_5_3></QID13_5_3>
		<QID13_4_1></QID13_4_1>
		<QID13_4_2></QID13_4_2>
		<QID13_4_3></QID13_4_3>
		<QID13_4_TEXT></QID13_4_TEXT>
		<QID6_1></QID6_1>
		<QID6_2>Click to write Scale point 2</QID6_2>
		<QID6_3>Click to write Scale point 1</QID6_3>
		<QID16_FIRST_CLICK>3.172</QID16_FIRST_CLICK>
		<QID16_LAST_CLICK>25.605</QID16_LAST_CLICK>
		<QID16_PAGE_SUBMIT>26.387</QID16_PAGE_SUBMIT>
		<QID16_CLICK_COUNT>9</QID16_CLICK_COUNT>
		<QID7_TEXT>foo</QID7_TEXT>
		<QID8_TEXT>bar</QID8_TEXT>
		<QID9_1>name</QID9_1>
		<QID9_2>email</QID9_2>
		<QID9_3>job role</QID9_3>
		<QID10_1>1</QID10_1>
		<QID10_2>2</QID10_2>
		<QID10_3>3</QID10_3>
		<QID17_0_GROUP_1>Item 1</QID17_0_GROUP_1>
		<QID17_0_GROUP_2>Item 2</QID17_0_GROUP_2>
		<QID17_0_GROUP_3></QID17_0_GROUP_3>
		<QID17_0_GROUP_4></QID17_0_GROUP_4>
		<QID17_0_GROUP_5></QID17_0_GROUP_5>
		<QID17_1_GROUP_1></QID17_1_GROUP_1>
		<QID17_1_GROUP_2></QID17_1_GROUP_2>
		<QID17_1_GROUP_3>Item 3</QID17_1_GROUP_3>
		<QID17_1_GROUP_4>Item 4</QID17_1_GROUP_4>
		<QID17_1_GROUP_5></QID17_1_GROUP_5>
		<QID17_2_GROUP_1></QID17_2_GROUP_1>
		<QID17_2_GROUP_2></QID17_2_GROUP_2>
		<QID17_2_GROUP_3></QID17_2_GROUP_3>
		<QID17_2_GROUP_4></QID17_2_GROUP_4>
		<QID17_2_GROUP_5></QID17_2_GROUP_5>
		<QID17_G0_1_RANK>2</QID17_G0_1_RANK>
		<QID17_G0_2_RANK>1</QID17_G0_2_RANK>
		<QID17_G0_3_RANK></QID17_G0_3_RANK>
		<QID17_G0_4_RANK></QID17_G0_4_RANK>
		<QID17_G0_5_RANK></QID17_G0_5_RANK>
		<QID17_G1_1_RANK></QID17_G1_1_RANK>
		<QID17_G1_2_RANK></QID17_G1_2_RANK>
		<QID17_G1_3_RANK>2</QID17_G1_3_RANK>
		<QID17_G1_4_RANK>1</QID17_G1_4_RANK>
		<QID17_G1_5_RANK></QID17_G1_5_RANK>
		<QID17_G2_1_RANK></QID17_G2_1_RANK>
		<QID17_G2_2_RANK></QID17_G2_2_RANK>
		<QID17_G2_3_RANK></QID17_G2_3_RANK>
		<QID17_G2_4_RANK></QID17_G2_4_RANK>
		<QID17_G2_5_RANK></QID17_G2_5_RANK>
		<QID17_5_TEXT></QID17_5_TEXT>
		<QID15_NPS_GROUP>Promoter</QID15_NPS_GROUP>
		<QID15>10</QID15>
		<Q_DataPolicyViolations></Q_DataPolicyViolations>
	</Response>
	<Response>
		<startDate>2019-08-20 12:51:41</startDate>
		<endDate>2019-08-20 12:52:04</endDate>
		<status>IP Address</status>
		<ipAddress>*******</ipAddress>
		<progress>33</progress>
		<duration>22</duration>
		<finished>False</finished>
		<recordedDate>2019-08-20 12:52:35</recordedDate>
		<_recordId>R_3MPTb9vwnCBmijR</_recordId>
		<recipientLastName>*******</recipientLastName>
		<recipientFirstName>*******</recipientFirstName>
		<recipientEmail>*******</recipientEmail>
		<externalDataReference>*******</externalDataReference>
		<locationLatitude>*******</locationLatitude>
		<locationLongitude>*******</locationLongitude>
		<distributionChannel>anonymous</distributionChannel>
		<userLanguage>EN</userLanguage>
		<QID1>Click to write Choice 2</QID1>
		<QID3>Click to write Choice 2</QID3>
		<QID4_1></QID4_1>
		<QID4_2>Click to write Choice 2</QID4_2>
		<QID4_3></QID4_3>
		<QID4_5></QID4_5>
		<QID4_6></QID4_6>
		<QID4_4></QID4_4>
		<QID4_5_TEXT></QID4_5_TEXT>
		<QID4_6_TEXT></QID4_6_TEXT>
		<QID11></QID11>
		<QID11_3_TEXT></QID11_3_TEXT>
		<QID5_1></QID5_1>
		<QID5_2></QID5_2>
		<QID5_3></QID5_3>
		<QID5_4></QID5_4>
		<QID5_4_TEXT></QID5_4_TEXT>
		<QID13_1_1></QID13_1_1>
		<QID13_1_2></QID13_1_2>
		<QID13_1_3></QID13_1_3>
		<QID13_2_1></QID13_2_1>
		<QID13_2_2></QID13_2_2>
		<QID13_2_3></QID13_2_3>
		<QID13_3_1></QID13_3_1>
		<QID13_3_2></QID13_3_2>
		<QID13_3_3></QID13_3_3>
		<QID13_5_1></QID13_5_1>
		<QID13_5_2></QID13_5_2>
		<QID13_5_3></QID13_5_3>
		<QID13_4_1></QID13_4_1>
		<QID13_4_2></QID13_4_2>
		<QID13_4_3></QID13_4_3>
		<QID13_4_TEXT></QID13_4_TEXT>
		<QID6_1></QID6_1>
		<QID6_2></QID6_2>
		<QID6_3></QID6_3>
		<QID16_FIRST_CLICK></QID16_FIRST_CLICK>
		<QID16_LAST_CLICK></QID16_LAST_CLICK>
		<QID16_PAGE_SUBMIT></QID16_PAGE_SUBMIT>
		<QID16_CLICK_COUNT></QID16_CLICK_COUNT>
		<QID7_TEXT></QID7_TEXT>
		<QID8_TEXT></QID8_TEXT>
		<QID9_1></QID9_1>
		<QID9_2></QID9_2>
		<QID9_3></QID9_3>
		<QID10_1></QID10_1>
		<QID10_2></QID10_2>
		<QID10_3></QID10_3>
		<QID17_0_GROUP_1></QID17_0_GROUP_1>
		<QID17_0_GROUP_2></QID17_0_GROUP_2>
		<QID17_0_GROUP_3></QID17_0_GROUP_3>
		<QID17_0_GROUP_4></QID17_0_GROUP_4>
		<QID17_0_GROUP_5></QID17_0_GROUP_5>
		<QID17_1_GROUP_1></QID17_1_GROUP_1>
		<QID17_1_GROUP_2></QID17_1_GROUP_2>
		<QID17_1_GROUP_3></QID17_1_GROUP_3>
		<QID17_1_GROUP_4></QID17_1_GROUP_4>
		<QID17_1_GROUP_5></QID17_1_GROUP_5>
		<QID17_2_GROUP_1></QID17_2_GROUP_1>
		<QID17_2_GROUP_2></QID17_2_GROUP_2>
		<QID17_2_GROUP_3></QID17_2_GROUP_3>
		<QID17_2_GROUP_4></QID17_2_GROUP_4>
		<QID17_2_GROUP_5></QID17_2_GROUP_5>
		<QID17_G0_1_RANK></QID17_G0_1_RANK>
		<QID17_G0_2_RANK></QID17_G0_2_RANK>
		<QID17_G0_3_RANK></QID17_G0_3_RANK>
		<QID17_G0_4_RANK></QID17_G0_4_RANK>
		<QID17_G0_5_RANK></QID17_G0_5_RANK>
		<QID17_G1_1_RANK></QID17_G1_1_RANK>
		<QID17_G1_2_RANK></QID17_G1_2_RANK>
		<QID17_G1_3_RANK></QID17_G1_3_RANK>
		<QID17_G1_4_RANK></QID17_G1_4_RANK>
		<QID17_G1_5_RANK></QID17_G1_5_RANK>
		<QID17_G2_1_RANK></QID17_G2_1_RANK>
		<QID17_G2_2_RANK></QID17_G2_2_RANK>
		<QID17_G2_3_RANK></QID17_G2_3_RANK>
		<QID17_G2_4_RANK></QID17_G2_4_RANK>
		<QID17_G2_5_RANK></QID17_G2_5_RANK>
		<QID17_5_TEXT></QID17_5_TEXT>
		<QID15_NPS_GROUP></QID15_NPS_GROUP>
		<QID15></QID15>
		<Q_DataPolicyViolations></Q_DataPolicyViolations>
	</Response>
</Responses>`

// spell-checker: enable
