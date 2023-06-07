package ptlist

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtListHappyPaths(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Europe/Athens", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Europe/Athens", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T220000Z",
					"20211101T220000Z",
					"20211102T220000Z",
					"20211103T220000Z",
					"20211104T220000Z",
					"20211105T220000Z",
					"20211106T220000Z",
					"20211107T220000Z",
					"20211108T220000Z",
					"20211109T220000Z",
					"20211110T220000Z",
					"20211111T220000Z",
					"20211112T220000Z",
					"20211113T220000Z",
					"20211114T220000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Europe/Athens", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T220000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T220000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Europe/Athens", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T220000Z",
					"20191231T220000Z",
					"20201231T220000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListUnhappyPaths(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
		expectedError  error
	}{
		{
			name:           "Hour test",
			input:          []string{"1h", "Europe/Aten", "", "20210715T123456Z"},
			expectedOutput: nil,
			expectedError:  errors.New("unknown time zone Europe/Aten"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.Error(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathSwedenTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Europe/Stockholm", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Europe/Stockholm", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T220000Z",
					"20211101T220000Z",
					"20211102T220000Z",
					"20211103T220000Z",
					"20211104T220000Z",
					"20211105T220000Z",
					"20211106T220000Z",
					"20211107T220000Z",
					"20211108T220000Z",
					"20211109T220000Z",
					"20211110T220000Z",
					"20211111T220000Z",
					"20211112T220000Z",
					"20211113T220000Z",
					"20211114T220000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Europe/Stockholm", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T220000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T220000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Europe/Stockholm", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T220000Z",
					"20191231T220000Z",
					"20201231T220000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}

func TestPtListHappyPathAbidjanTZ(t *testing.T) {
	testcases := []struct {
		name           string
		input          []string
		expectedOutput *PtListResponse
	}{
		{
			name:  "Hour test",
			input: []string{"1h", "Africa/Abidjan", "20210714T204603Z", "20210715T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210714T210000Z",
					"20210714T220000Z",
					"20210714T230000Z",
					"20210715T000000Z",
					"20210715T010000Z",
					"20210715T020000Z",
					"20210715T030000Z",
					"20210715T040000Z",
					"20210715T050000Z",
					"20210715T060000Z",
					"20210715T070000Z",
					"20210715T080000Z",
					"20210715T090000Z",
					"20210715T100000Z",
					"20210715T110000Z",
					"20210715T120000Z"},
			},
		},
		{
			name:  "Day test",
			input: []string{"1d", "Africa/Abidjan", "20211010T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20211010T210000Z",
					"20211011T210000Z",
					"20211012T210000Z",
					"20211013T210000Z",
					"20211014T210000Z",
					"20211015T210000Z",
					"20211016T210000Z",
					"20211017T210000Z",
					"20211018T210000Z",
					"20211019T210000Z",
					"20211020T210000Z",
					"20211021T210000Z",
					"20211022T210000Z",
					"20211023T210000Z",
					"20211024T210000Z",
					"20211025T210000Z",
					"20211026T210000Z",
					"20211027T210000Z",
					"20211028T210000Z",
					"20211029T210000Z",
					"20211030T210000Z",
					"20211031T210000Z",
					"20211101T210000Z",
					"20211102T210000Z",
					"20211103T210000Z",
					"20211104T210000Z",
					"20211105T210000Z",
					"20211106T210000Z",
					"20211107T210000Z",
					"20211108T210000Z",
					"20211109T210000Z",
					"20211110T210000Z",
					"20211111T210000Z",
					"20211112T210000Z",
					"20211113T210000Z",
					"20211114T210000Z"},
			},
		},
		{
			name:  "Month test",
			input: []string{"1mo", "Africa/Abidjan", "20210214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20210228T210000Z",
					"20210331T210000Z",
					"20210430T210000Z",
					"20210531T210000Z",
					"20210630T210000Z",
					"20210731T210000Z",
					"20210831T210000Z",
					"20210930T210000Z",
					"20211031T210000Z"},
			},
		},
		{
			name:  "Year test",
			input: []string{"1y", "Africa/Abidjan", "20180214T204603Z", "20211115T123456Z"},
			expectedOutput: &PtListResponse{
				[]string{
					"20181231T210000Z",
					"20191231T210000Z",
					"20201231T210000Z"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			srv := NewService()

			ptlist, err := srv.GetPtList(context.Background(), tc.input[0], tc.input[1], tc.input[2], tc.input[3])
			if err != nil {
				assert.NoError(t, errors.New(err.Desc))
			}
			require.Equal(t, tc.expectedOutput, ptlist)
		})
	}
}
